package main

import (
	"github.com/mitoteam/goapp"
	"github.com/mitoteam/mt-tgadmin/app"
)

func main() {
	//create app and set it up
	app.App = goapp.NewAppBase(app.Settings)

	app.App.AppName = "mt-tgadmin"
	app.App.ExecutableName = "mt-tgadmin"
	app.App.LongDescription = `Simple telegram bot Web GUI to send messages to groups.

	Copyright: MiTo Team, https://mito-team.com`

	app.App.AppSettingsFilename = ".bot.yml"

	//use default gin router
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
