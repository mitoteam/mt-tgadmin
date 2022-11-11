// Package app contains main mt-tgadmin functionality.
package app

type Settings struct {
	BotToken string `yaml:"bot_token" yaml_comment:"Bot authorization token"`

	GuiCookieSecretKey string `yaml:"gui_cookie_secret_key" yaml_comment:"At least 32 chars"`
	GuiPassword        string `yaml:"gui_password" yaml_comment:"GUI access password"`
	GuiHostname        string `yaml:"gui_hostname" yaml_comment:"Web GUI hostname"`
	GuiPort            uint16 `yaml:"gui_port" yaml_comment:"Web GUI port number"`
}
