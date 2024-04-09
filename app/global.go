// Package app contains main mt-tgadmin functionality.
package app

var Global struct {
	BotInfo  string
	ChatInfo string
}

func init() {
	//defaults
	Global.BotInfo = "[undefined]"
	Global.ChatInfo = "[undefined]"
}
