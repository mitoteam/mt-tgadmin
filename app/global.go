// Package app contains main mt-tgadmin functionality.
package app

import "github.com/mitoteam/goappbase"

var (
	Global struct {
		BotInfo  string
		ChatInfo string
	}

	App      *goappbase.AppBase
	Settings *AppSettingsType
)

func init() {
	//defaults
	Global.BotInfo = "[undefined]"
	Global.ChatInfo = "[undefined]"
}
