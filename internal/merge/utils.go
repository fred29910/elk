package merge

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// parseFile parses a Go source file and returns its AST.
func parseFile(filename string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}
	return fset, node, nil
}

// parseBytes parses a Go source bytes and returns its AST.
func parseBytes(bs []byte) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", string(bs), parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}
	return fset, node, nil
}
