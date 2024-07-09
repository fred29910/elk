
build:
	go build --trimpath -o bin/ ./cmd/...

clean:
	rm -rf bin
.PHONY: build clean