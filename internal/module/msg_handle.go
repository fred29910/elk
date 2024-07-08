package module

import (
	"bytes"
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/masseelch/elk/internal/generator"
	"github.com/masseelch/elk/internal/parser"
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
	//m.MSGValue = make(map[string]int32)
	//m.ctx
	return m
}

func (m *Module) Name() string { return "msg_handle" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}
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
	m.OverwriteCustomFile(
		"./tmp/report.txt",
		buf.String(),
		0644,
	)

	return m.Artifacts()
}

func (m *Module) getPackInfo(pkg pgs.Package) error {
	// common msg
	var commonMsgs []string
	msgs := make(map[string]*parser.Msg)
	for _, file := range pkg.Files() {
		if commonMsgs == nil {
			commonMsgs = parser.CommonsParser(file.SourceCodeInfo().LeadingComments())
		}
		for _, message := range file.AllMessages() {
			msgDf := parser.MsgParser(message.SourceCodeInfo().LeadingComments())
			if msgDf == nil {
				continue
			}
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
		Name:    pkg.ProtoName().String(),
		Imports: nil,
		Coms:    csgs,
		Cms:     msgs,
		OutDir:  "",
	}
	return nil
}

func (m *Module) getMsgs(pkgs map[string]pgs.Package) error {
	var cnmGoPackage string
	for s, p := range pkgs {
		m.Logf("package name is %s", s)
		for _, f := range p.Files() {

			cnmGoPackage = f.Descriptor().GetOptions().GetGoPackage()
			for _, eVaule := range f.Enums() {
				if eVaule.Name() == msgName {
					m.MSGValue.Path = s
					m.MSGValue.Msgs = make(map[string]int32, len(eVaule.Values()))
					for _, value := range eVaule.Values() {
						m.MSGValue.Msgs[value.Name().String()] = value.Value()
					}
					break
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
	m.MSGValue.Pkg = cnmGoPackage
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
