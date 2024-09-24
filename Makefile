.PHONY: build gen clean
build:
	go build --trimpath -o build/  ./...

gen:
	swag init -g cmd/coderx/main.go


clean:
	rm -rf internal/gen
	rm -rf build