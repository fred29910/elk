package main

import (
	"fmt"
	"github.com/masseelch/elk/internal/merge"
)

func main() {

	file1 := "s/badge.go"
	file2 := "ss/badge.go"

	tx, err := merge.NewCompare(file1, file2)
	if err != nil {
		panic(err)
	}
	tar, err := merge.FormatNode(tx)
	if err != nil {
		panic(err)
	}
	fmt.Println(tar)
}
