
build:
	go build --trimpath -o bin/ ./cmd/...

clean:
	rm -rf bin


init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/lyft/protoc-gen-star/protoc-gen-debug@latest

.PHONY: build clean