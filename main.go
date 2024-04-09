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

	app.Settings = &app.AppSettingsType{
		GuiPassword: "mitoteam",
	}
	app.Settings.WebserverPort = 15080

	app.App = goappbase.NewAppBase(app.Settings)

	app.App.AppName = "mt-tgadmin"
	app.App.ExecutableName = "mt-tgadmin"
	app.App.LongDescription = `simple telegram bot Web GUI based manager to send messages to groups.

	Copyright: MiTo Team, https://mito-team.com`

	app.App.AppSettingsFilename = ".bot.yml"

	app.App.BuildWebRouterF = app.BuildWebRouter

	app.App.PreRunF = func() error {
		err := app.InitApi()

		return err
	}

	app.App.Run()
}
