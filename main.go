package main

import (
	"embed"
	"log"
	"mt-tgadmin/app"
	"mt-tgadmin/cmd"
)

// embedded web assets
//
//go:embed assets/*.min.js assets/*.css
var webAssets embed.FS

//go:embed assets/index.html
var webIndexHtml string

func main() {
	app.WebAssets = &webAssets
	app.WebIndexHtml = &webIndexHtml

	//cli application - we just let cobra to do it job
	if err := cmd.Root().Execute(); err != nil {
		log.Fatalln(err)
	}
}
