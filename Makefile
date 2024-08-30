.PHONY: build 
build:
	go build --trimpath -o build/  ./cmd/... 

.PHONY: gendesign
gendesign:
	goa gen github.com/cham/elk/design
	goa example github.com/cham/elk/design  -o pkg/servicex/
	cd web/ui && yarn genapi