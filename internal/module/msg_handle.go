package module

import (
	"bytes"
	"fmt"
	pgs "github.com/lyft/protoc-gen-star/v2"
)

type Module struct {
	*pgs.ModuleBase
	//ctx pgsgo.Context
	MSGValue map[string]int32
}

const msgName = "MSG"

func New() pgs.Module { return &Module{&pgs.ModuleBase{}, make(map[string]int32)} }

func (m *Module) Name() string { return "msg_handle" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for s, f := range targets {
		goPkg := f.Descriptor().GetOptions().GetGoPackage()
		protoPkg := f.Package().ProtoName().String()
		m.Logf("Protobuf package: %s maps to Go package: %s", protoPkg, goPkg)
		for eIndex, eVaule := range f.AllEnums() {
			if eVaule.Name() == msgName {
				for _, value := range eVaule.Values() {
					m.MSGValue[value.Name().String()] = value.Value()
					fmt.Fprintf(buf, "---%v %v %v --- \n", eVaule.Name(), value.Name().String(), value.Value())
				}
			}

			fmt.Fprintf(buf, "%03d. %v\n", eIndex, eVaule.Name())
			m.Debug(fmt.Sprintf("enums all is pxv %v", eVaule))
		}

		m.Push(f.Name().String()).Debug("reporting")

		fmt.Fprintf(buf, "---%v %v --- \n", s, f.Package().ProtoName().LowerSnakeCase())

		for i, msg := range f.AllMessages() {
			m.Debug(msg.SourceCodeInfo().LeadingComments())
			fmt.Fprintf(buf, "%03d. %v\n", i, msg.Name())
			msg.Descriptor()
			m.Debug(fmt.Sprintf("enums all is pxv %v", msg))
		}

		m.Pop()

	}
	//for s, p := range pkgs {
	//	fmt.Fprintf(buf, "%v --- %v\n", s, p.ProtoName().String())
	//	for i, file := range p.Files() {
	//
	//	}
	//}

	m.OverwriteCustomFile(
		"./tmp/report.txt",
		buf.String(),
		0644,
	)

	return m.Artifacts()
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
