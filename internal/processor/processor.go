package processor

import (
	"fmt"
	"time"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// New returns a PostProcessor that adds a copyright comment to the top
// of all generated files.
func New(owner string) pgs.PostProcessor { return copyrightPostProcessor{owner} }

type copyrightPostProcessor struct {
	owner string
}

// Match returns true only for Custom and Generated files (including templates).
func (cpp copyrightPostProcessor) Match(a pgs.Artifact) bool {
	switch a := a.(type) {
	case pgs.GeneratorFile, pgs.GeneratorTemplateFile,
		pgs.CustomFile, pgs.CustomTemplateFile:
		return true
	default:
		if a != nil {

		}
		return false
	}
}

// Process attaches the copyright header to the top of the input bytes
func (cpp copyrightPostProcessor) Process(in []byte) (out []byte, err error) {
	cmt := fmt.Sprintf("// Copyright © %d %s. All rights reserved\n",
		time.Now().Year(),
		cpp.owner)

	return append([]byte(cmt), in...), nil
}
