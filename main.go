package main

import (
	"github.com/mitoteam/goappbase"
	"github.com/mitoteam/mt-tgadmin/app"
)

func main() {
	//default settings
	app.Settings = &app.AppSettingsType{
		GuiPassword: "mitoteam",
	}
	app.Settings.WebserverPort = 15080

	//create app and set it up
	app.App = goappbase.NewAppBase(app.Settings)

	app.App.AppName = "mt-tgadmin"
	app.App.ExecutableName = "mt-tgadmin"
	app.App.LongDescription = `simple telegram bot Web GUI based manager to send messages to groups.

	Copyright: MiTo Team, https://mito-team.com`

	app.App.AppSettingsFilename = ".bot.yml"

	//router
	app.App.BuildWebRouterF = app.BuildWebRouter

	//API
	app.BuildWebApiRouter(app.App)

	//initialization
	app.App.PreRunF = func() error {
		err := app.InitTgApi()

		return err
	}

	//global state default values
	app.App.Global["BotInfo"] = "[undefined]"
	app.App.Global["ChatInfo"] = "[undefined]"

	//do all the job
	app.App.Run()
}
