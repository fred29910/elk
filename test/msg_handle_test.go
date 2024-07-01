package test

import (
	"bytes"
	"os"
	"testing"

	pgs "github.com/lyft/protoc-gen-star/v2"
	"github.com/masseelch/elk/internal/module"
	"github.com/spf13/afero"
)

func TestModule(t *testing.T) {
	req, err := os.Open("./code_generator_request.pb.bin")
	if err != nil {
		t.Fatal(err)
	}

	fs := afero.NewMemMapFs()
	res := &bytes.Buffer{}

	pgs.Init(
		pgs.ProtocInput(req),  // use the pre-generated request
		pgs.ProtocOutput(res), // capture CodeGeneratorResponse
		pgs.FileSystem(fs),    // capture any custom files written directly to disk
	).RegisterModule(module.New()).Render()

	// check res and the fs for output
}
