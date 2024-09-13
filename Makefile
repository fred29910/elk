.PHONY: build gen clean
build:
	go build --trimpath -o build/  ./...

gen:
	goa gen github.com/cham/elk/design -o internal


clean:
	rm -rf internal/gen
	rm -rf build