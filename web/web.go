package web

import (
	"embed"
	iofs "io/fs"
	"net/http"
	"strings"
)

//go:embed ui/dist
var UiDs embed.FS

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {

	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if p == "" {
			p = "/index.html"
		}
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	sub, err := iofs.Sub(UiDs, root)
	if err != nil {
		panic(err)
	}
	fs := http.FS(sub)
	return &binaryFileSystem{
		fs,
	}
}
