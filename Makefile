.PHONY: build gen clean
build: gen
	go build --trimpath -ldflags "-s -w"  -o build/  ./...

gen:
	swag init --parseVendor=true --parseDependency=true -g cmd/coderx/main.go


clean:
	rm -rf internal/gen
	rm -rf build