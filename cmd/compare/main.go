package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {

	file1 := "s/badge.go"
	file2 := "ss/badge.go"

	_, f1, err := parseFile(file1)
	if err != nil {
		panic(err)
	}
	_, f2, err := parseFile(file2)
	if err != nil {
		panic(err)
	}
	print(f1, f2)
}

// parseFile parses a Go source file and returns its AST.
func parseFile(filename string) (*token.FileSet, *ast.File, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		return nil, nil, err
	}
	return fset, node, nil
}
