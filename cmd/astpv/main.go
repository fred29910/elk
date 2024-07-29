package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	js.Global().Set("add", js.FuncOf(add))

	<-c
}

func add(this js.Value, p []js.Value) interface{} {
	a := p[0].Int()
	b := p[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}
