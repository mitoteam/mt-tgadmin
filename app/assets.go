package app

import (
	"embed"
	"io/fs"
)

// embedded assets and templates
//
//go:embed assets/*.min.js assets/vue.global.prod.js
//go:embed assets/*.css
//go:embed assets/favicon.ico
//go:embed templates/*.html
var embedFS embed.FS

var webAssetsFS fs.FS
var templatesFS fs.FS

func init() {
	//prepare FS for subdirectory "/assets"
	webAssetsFS, _ = fs.Sub(embedFS, "assets")

	//prepare FS for subdirectory "/templates"
	templatesFS, _ = fs.Sub(embedFS, "templates")
}
