// Package app contains main mt-tgadmin functionality.
package app

type Settings struct {
	BotName string `yaml:"bot_name" yaml_comment:"Bot name to show in GUI"`

	GuiPassword string `yaml:"gui_password" yaml_comment:"GUI access password"`
	GuiHostname string `yaml:"gui_hostname" yaml_comment:"Web GUI hostname"`
	GuiPort     uint16 `yaml:"gui_port" yaml_comment:"Web GUI port number"`
}
