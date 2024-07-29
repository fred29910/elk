// compile: GOOS=js GOARCH=wasm
package main

import (
	"github.com/masseelch/elk/internal/astextract"
	"github.com/masseelch/elk/internal/ov"
	"github.com/masseelch/elk/internal/utils"
	"github.com/masseelch/elk/internal/utils/res"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")

	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("parser", js.FuncOf(parserOriginRange))

	<-c
}

func add(this js.Value, p []js.Value) interface{} {
	a := p[0].Int()
	b := p[1].Int()
	sum := a + b
	return js.ValueOf(sum)
}

func parserOriginRange(this js.Value, args []js.Value) interface{} {

	in := args[0].String()

	req, err := utils.UnmarshalJSON[ov.ParserReq]([]byte(in))
	if err != nil {
		return res.FailRes(err.Error(), 1)
	}
	outer, err := astextract.Parse(req.Content)
	if err != nil {
		return res.FailRes(err.Error(), 2)
	}

	return res.SuccessRes(outer)
}
