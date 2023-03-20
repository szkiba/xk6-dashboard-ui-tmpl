package ui

import (
	"embed"
	"io/fs"
	"net/http"
)

const (
	dir    = "ui"
	prefix = "/ui/"
)

//go:embed ui
var filesystem embed.FS

func init() {
	subfs, err := fs.Sub(filesystem, dir)
	if err != nil {
		panic(err)
	}

	http.DefaultServeMux.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.FS(subfs))))
}
