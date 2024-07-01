package module

import (
	"bytes"
	"fmt"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

type Module struct {
	*pgs.ModuleBase
	// ctx pgsgo.Context
}

func New() pgs.Module { return &Module{&pgs.ModuleBase{}} }

func (m *Module) Name() string { return "msg_handle" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}

	for _, f := range targets {
		m.Push(f.Name().String()).Debug("reporting")

		fmt.Fprintf(buf, "--- %v --- \n", f.Name())

		for i, msg := range f.AllMessages() {
			m.Debug(msg.SourceCodeInfo().LeadingComments())
			fmt.Fprintf(buf, "%03d. %v\n", i, msg.Name())
			msg.Descriptor()

			m.Debug(fmt.Sprintf("enums all is pxv %v", msg))
		}

		m.Pop()

		for eIndex, eVaule := range f.AllEnums() {
			fmt.Fprintf(buf, "%03d. %v\n", eIndex, eVaule.Name())
			m.Debug(fmt.Sprintf("enums all is pxv %v", eVaule))
		}
	}

	m.OverwriteCustomFile(
		"./tmp/report.txt",
		buf.String(),
		0644,
	)

	return m.Artifacts()
}
