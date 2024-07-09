package module

import (
	"fmt"
	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/masseelch/elk/internal/generator"
	"github.com/masseelch/elk/internal/parser"
	"path/filepath"
)

type Module struct {
	*pgs.ModuleBase
	//ctx      pgsgo.Context
	MSGValue generator.MsgDefine
	pkgs     map[string]generator.Package
}

const msgName = "MSG"

func New() pgs.Module {
	m := &Module{ModuleBase: &pgs.ModuleBase{}}
	m.pkgs = make(map[string]generator.Package)
	//m.ctx
	return m
}

func (m *Module) Name() string { return "msg_handle" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	//buf := &bytes.Buffer{}
	for s, s2 := range m.Parameters() {
		m.Logf("%v ---> %v", s, s2)
	}

	err := m.getMsgs(pkgs)
	if err != nil {
		m.Logf("parser error %s ", err.Error())
		return nil
	}

	for s, p := range pkgs {
		if s == m.MSGValue.Path {
			continue
		}
		err = m.getPackInfo(p)
		if err != nil {
			m.Logf("parser error %s ", err.Error())
			return nil
		}
	}

	for s, g := range m.pkgs {
		m.Debugf("%v package out dir sx %s", s, m.Parameters().Str(s))
		str, err := generator.Gen(m.MSGValue, g)
		if err != nil {
			m.Logf("cat not gen file %v", err.Error())
			return nil
		}

		outpath := filepath.Join(m.OutputPath(), m.Parameters().Str(s), "msg.go")
		m.Debugf("%v package out dir ", outpath)
		m.OverwriteCustomFile(
			outpath,
			str,
			0644,
		)
	}

	return m.Artifacts()
}

func (m *Module) getPackInfo(pkg pgs.Package) error {
	// common msg
	var commonMsgs []string
	var cnmGoPackage string
	msgs := make(map[string]*parser.Msg)
	for _, file := range pkg.Files() {
		cnmGoPackage = file.Descriptor().GetOptions().GetGoPackage()
		if commonMsgs == nil {
			// 获取文件描述符和源代码信息
			fd := file.Descriptor()
			locations := fd.GetSourceCodeInfo().GetLocation()
			// 遍历源代码位置，查找 package 声明的注释
			for _, loc := range locations {
				if len(loc.Path) > 0 && loc.Path[0] == int32(2) {
					if loc.LeadingComments != nil {
						commonMsgs = parser.CommonsParser(loc.GetLeadingComments())
						m.Logf("%v common msg %v", pkg.ProtoName().String(), commonMsgs)
					}
				}
			}
		}
		for _, message := range file.AllMessages() {
			msgDf := parser.MsgParser(message.SourceCodeInfo().LeadingComments())
			if msgDf == nil {
				continue
			}
			m.Debugf("msg name %v", message.Name().UpperCamelCase())
			msgDf.Name = message.Name().UpperCamelCase().String()
			if _, ok := m.MSGValue.Msgs[msgDf.ID]; !ok {
				return fmt.Errorf("cat not get msg %s", msgDf.ID)
			}
			msgs[msgDf.ID] = msgDf
		}
	}
	csgs := make(map[string]*parser.Msg)
	for _, msg := range commonMsgs {
		mdf, ok := m.MSGValue.Cms[msg]
		if !ok {
			return fmt.Errorf("cat not get msg %s", mdf.ID)
		}
		csgs[msg] = mdf
	}
	m.pkgs[pkg.ProtoName().String()] = generator.Package{
		Name:          pkg.ProtoName().String(),
		Imports:       nil,
		Coms:          csgs,
		Cms:           msgs,
		GopackageName: cnmGoPackage,
		OutDir:        "",
	}
	return nil
}

func (m *Module) getMsgs(pkgs map[string]pgs.Package) error {
outerLoop:
	for s, p := range pkgs {
		m.Debugf("package name is %s", s)
		for _, f := range p.Files() {
			for _, eVaule := range f.Enums() {
				if eVaule.Name() == msgName {
					m.MSGValue.Path = s
					m.MSGValue.Pkg = f.Descriptor().GetOptions().GetGoPackage()
					m.MSGValue.Msgs = make(map[string]int32, len(eVaule.Values()))
					for _, value := range eVaule.Values() {
						m.MSGValue.Msgs[value.Name().String()] = value.Value()
					}
					break outerLoop
				}
			}
		}
	}

	cms := make(map[string]*parser.Msg)
	// common msg
	cnmPkg := pkgs[m.MSGValue.Path]
	for _, file := range cnmPkg.Files() {

		for _, message := range file.AllMessages() {
			msgDf := parser.MsgParser(message.SourceCodeInfo().LeadingComments())
			if msgDf == nil {
				continue
			}
			if _, ok := m.MSGValue.Msgs[msgDf.ID]; !ok {
				return fmt.Errorf("cat not get msg %s", msgDf.ID)
			}
			cms[msgDf.ID] = msgDf
		}
	}
	m.MSGValue.Cms = cms

	return nil
}

type Generator struct {
	Common  string              // 公共包 gopath
	Service map[string]*Service //具体服务的代码, key gopath
}

type Service struct {
	Name   string            // handle interface name
	Common map[string]string // 公共服务Key与msg
	Pmv    map[string]string // 非公共的消息与绑定

}
