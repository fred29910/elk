package main

import (
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
)

func main() {
	code := `package a

func main(){
	var a int    // foo
	var b string // bar
}
`
	f, err := decorator.Parse(code)
	if err != nil {
		panic(err)
	}

	list := f.Decls[0].(*dst.FuncDecl).Body.List
	list[0], list[1] = list[1], list[0]

	if err := decorator.Print(f); err != nil {
		panic(err)
	}
}
