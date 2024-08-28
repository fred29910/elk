.PHONY: build 
build:
	go build --trimpath -o build/ ./pusher/cmd/pusher
	go build --trimpath -o build/ ./consumer/cmd/consumer