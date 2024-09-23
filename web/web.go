package web

import (
	"embed"
	"net/http"
	"path"
	"strings"
)

//go:embed ui/dist
var UiDs embed.FS

// PrefixFileSystem wraps an existing http.FileSystem to remove a prefix from the file path.
type PrefixFileSystem struct {
	fs     http.FileSystem
	prefix string
}

func (pfs *PrefixFileSystem) Open(name string) (http.File, error) {
	// Remove the prefix from the file path
	name = strings.TrimPrefix(name, pfs.prefix)
	return pfs.fs.Open(name)
}

// NewPrefixFileSystem creates a new PrefixFileSystem
func NewPrefixFileSystem(fs embed.FS, prefix string) http.FileSystem {
	// subFS, err := iofs.Sub(fs, prefix)
	// if err != nil {
	// 	panic(err)
	// }
	return &PrefixFileSystem{
		fs:     http.FS(fs),
		prefix: prefix,
	}
}

// GetContentType returns the content type based on the file extension
func GetContentType(filePath string) string {
	ext := path.Ext(filePath)
	switch ext {
	case ".js":
		return "application/javascript"
	case ".css":
		return "text/css"
	case ".html":
		return "text/html"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
