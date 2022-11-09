// Package app contains main mt-tgadmin functionality.
package app

import (
	"runtime"
)

var BuildVersion = "DEV"
var BuildCommit = "DEV"

const DefaultSettingsFilename = ".bot.yml"

var Global struct {
	AppName   string
	Version   string
	Commit    string
	BuiltWith string

	SettingsFilename string
}

func init() {
	Global.AppName = "mt-tgadmin"
	Global.Version = BuildVersion
	Global.Commit = BuildCommit
	Global.BuiltWith = runtime.Version()
	Global.SettingsFilename = DefaultSettingsFilename
}
