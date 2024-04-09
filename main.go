package main

import (
	"embed"

	"github.com/mitoteam/goappbase"
	"github.com/mitoteam/mt-tgadmin/app"
)

// embedded web assets
//
//go:embed assets/*.min.js assets/*.css assets/favicon.ico
var webAssets embed.FS

//go:embed assets/index.html
var webIndexHtml string

func main() {
	app.WebAssets = &webAssets
	app.WebIndexHtml = &webIndexHtml

	settings := &app.AppSettingsType{
		GuiPassword: "mitoteam",
	}
	settings.WebserverPort = 15080

	application := goappbase.NewAppBase(settings)

	application.AppName = "mt-tgadmin"
	application.ExecutableName = "mt-tgadmin"
	application.LongDescription = `simple telegram bot Web GUI based manager to send messages to groups.

	Copyright: MiTo Team, https://mito-team.com`

	application.AppSettingsFilename = ".bot.yml"

	application.BuildWebRouterF = app.BuildWebRouter

	application.PreRunF = func() error {
		err := app.InitApi(application)

		return err
	}

	application.Run()
}
