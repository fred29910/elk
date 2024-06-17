package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
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

// compareASTs compares two AST nodes and prints the differences.
func compareASTs(node1, node2 interface{}) {
	if !reflect.DeepEqual(node1, node2) {
		fmt.Println("ASTs are different")
	} else {
		fmt.Println("ASTs are identical")
	}
}

func main() {

	file1 := "s/badge.go"
	file2 := "ss/badge.go"

	_, node1, err1 := parseFile(file1)
	if err1 != nil {
		fmt.Printf("Error parsing %s: %v\n", file1, err1)
		return
	}

	_, node2, err2 := parseFile(file2)
	if err2 != nil {
		fmt.Printf("Error parsing %s: %v\n", file2, err2)
		return
	}

	compareASTs(node1, node2)
}
