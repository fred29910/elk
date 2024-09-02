.PHONY: build 
build:
	go build --trimpath -o build/  main.go


.PHONY: trace
trace:
	go tool trace trace.out