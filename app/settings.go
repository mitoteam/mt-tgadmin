// Package app contains main mt-tgadmin functionality.
package app

type Settings struct {
	BotName string `yaml:"bot_name" yaml_comment:"Bot name to show in GUI"`

	Password string `yaml:"password" yaml_comment:"GUI access password"`
}
