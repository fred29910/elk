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

// extractStructsAndFuncs extracts struct and function declarations from the AST.
func extractStructsAndFuncs(file *ast.File) (structs map[string]*ast.StructType, funcs map[string]*ast.FuncDecl) {
	structs = make(map[string]*ast.StructType)
	funcs = make(map[string]*ast.FuncDecl)

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if structType, ok := s.Type.(*ast.StructType); ok {
						structs[s.Name.Name] = structType
					}
				}
			}
		case *ast.FuncDecl:
			funcs[d.Name.Name] = d
		}
	}

	return structs, funcs
}

// compareStructs compares two maps of struct declarations.
func compareStructs(structs1, structs2 map[string]*ast.StructType) {
	for name, struct1 := range structs1 {
		if struct2, ok := structs2[name]; ok {
			if !reflect.DeepEqual(struct1, struct2) {
				fmt.Printf("Struct %s is different\n", name)
			} else {
				fmt.Printf("Struct %s is identical\n", name)
			}
		} else {
			fmt.Printf("Struct %s is only in the first file\n", name)
		}
	}

	for name := range structs2 {
		if _, ok := structs1[name]; !ok {
			fmt.Printf("Struct %s is only in the second file\n", name)
		}
	}
}

// compareFuncs compares two maps of function declarations.
func compareFuncs(funcs1, funcs2 map[string]*ast.FuncDecl) {
	for name, func1 := range funcs1 {
		if func2, ok := funcs2[name]; ok {
			if !reflect.DeepEqual(func1, func2) {
				fmt.Printf("Function %s is different\n", name)
			} else {
				fmt.Printf("Function %s is identical\n", name)
			}
		} else {
			fmt.Printf("Function %s is only in the first file\n", name)
		}
	}

	for name := range funcs2 {
		if _, ok := funcs1[name]; !ok {
			fmt.Printf("Function %s is only in the second file\n", name)
		}
	}
}
