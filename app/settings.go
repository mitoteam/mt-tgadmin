// Package app contains main mt-tgadmin functionality.
package app

import "github.com/mitoteam/goappbase"

type AppSettingsType struct {
	goappbase.AppSettingsBase `yaml:",inline"`

	BotToken  string `yaml:"bot_token" yaml_comment:"Bot authorization token"`
	BotChatID int64  `yaml:"bot_chat_id" yaml_comment:"'chat_id' int64 value"`

	GuiPassword string `yaml:"gui_password" yaml_comment:"GUI access password"`
}
