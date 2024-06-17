package main

import (
	"fmt"
)

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

	structs1, funcs1 := extractStructsAndFuncs(node1)
	structs2, funcs2 := extractStructsAndFuncs(node2)

	fmt.Println("Comparing Structs:")
	compareStructs(structs1, structs2)

	fmt.Println("Comparing Functions:")
	compareFuncs(funcs1, funcs2)
}
