// Package app contains main mt-tgadmin functionality.
package app

type Settings struct {
	BotToken  string `yaml:"bot_token" yaml_comment:"Bot authorization token"`
	BotChatID int64  `yaml:"bot_chat_id" yaml_comment:"'chat_id' int64 value"`

	GuiCookieSecretKey string `yaml:"gui_cookie_secret_key" yaml_comment:"At least 32 chars"`
	GuiPassword        string `yaml:"gui_password" yaml_comment:"GUI access password"`
	GuiHostname        string `yaml:"gui_hostname" yaml_comment:"Web GUI hostname"`
	GuiPort            uint16 `yaml:"gui_port" yaml_comment:"Web GUI port number"`

	ServiceName  string `yaml:"service_name" yaml_comment:"Service name for 'install' command"`
	ServiceUser  string `yaml:"service_user" yaml_comment:"User for 'install' command"`
	ServiceGroup string `yaml:"service_group" yaml_comment:"Group for 'install' command"`
}
