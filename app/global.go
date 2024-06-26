// Package app contains main mt-tgadmin functionality.
package app

import (
	"github.com/mitoteam/goappbase"
)

type AppSettingsType struct {
	goappbase.AppSettingsBase `yaml:",inline"`

	BotToken  string `yaml:"bot_token" yaml_comment:"Bot authorization token"`
	BotChatID int64  `yaml:"bot_chat_id" yaml_comment:"Telegram 'chat_id' int64 value"`

	GuiPassword string `yaml:"gui_password" yaml_comment:"GUI access password"`
}

var (
	App      *goappbase.AppBase
	Settings *AppSettingsType
)

func init() {
	//default settings
	Settings = &AppSettingsType{
		GuiPassword: "mitoteam",
	}

	//default values for goappbase.AppSettingsBase options
	Settings.WebserverPort = 15080
}
