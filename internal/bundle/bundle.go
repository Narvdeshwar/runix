package bundle

import (
	"embed"
	"io/fs"
	"net/http"
)

// UIAssets holds the bundled React UI files
//
//go:embed all:dist
var UIAssets embed.FS

// GetUIServer returns an http.Handler that serves the bundled UI
func GetUIServer() http.Handler {
	sub, err := fs.Sub(UIAssets, "dist")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(sub))
}
