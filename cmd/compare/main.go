package main

import (
	"github.com/masseelch/elk/internal/ast_compare"
)

func main() {

	file1 := "s/badge.go"
	file2 := "ss/badge.go"

	comer, err := ast_compare.NewCompare(file1, file2)
	if err != nil {
		panic(err)
	}
	err = comer.Compare()
	if err != nil {
		panic(err)
	}
}
