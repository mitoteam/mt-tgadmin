package app

import (
	"embed"
	"io/fs"
)

// embedded web assets
//
//go:embed assets/*.min.js assets/*.css assets/favicon.ico
//go:embed templates/*.html
var embedFS embed.FS

var webAssetsFS fs.FS
var templatesFS fs.FS

func init() {
	//prepare FS for subdirectory "/assets"
	webAssetsFS, _ = fs.Sub(embedFS, "assets")
	templatesFS, _ = fs.Sub(embedFS, "templates")
}
