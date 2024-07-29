.PHONY: build 
build:
	GOOS=js GOARCH=wasm go build --trimpath -o build/main.wasm ./cmd/astpv/main.go