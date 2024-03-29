// Package cmd provides CLI commands, flags and arguments handling.
// spf13/cobra based.
package cmd

import (
	"log"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/mitoteam/mttools"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     app.Global.AppName,
	Version: app.Global.Version,
	Long: app.Global.AppName + ` - simple telegram bot to send messages to group.

Copyright: MiTo Team, https://mito-team.com`,

	//disable 'completion' subcommand
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},

	Run: func(cmd *cobra.Command, args []string) {
		//show help if no subcommand given
		cmd.Help()
	},

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if mttools.IsFileExists(app.Global.SettingsFilename) {
			if err := mttools.LoadYamlSettingFromFile(app.Global.SettingsFilename, &app.Global.Settings); err != nil {
				return err
			}

			//mttools.PrintYamlSettings(app.Global.Settings)
		} else {
			if cmd.Name() != "init" && cmd.Name() != "version" {
				log.Fatalln(
					"No " + app.DefaultSettingsFilename + " file found. Please create one or use `" +
						app.Global.AppName + " init` command.",
				)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&app.Global.SettingsFilename,
		"settings",
		app.Global.SettingsFilename,
		"Filename or full path to bot settings file.",
	)
}

func Root() *cobra.Command {
	return rootCmd
}

// CallParentPreRun helper function calls parent command's PersistentPreRun
// or PersistentPreRunE hooks if they are defined.
func CallParentPreRun(cmd *cobra.Command, args []string) error {
	parent := cmd.Parent()

	if parent == nil {
		return nil
	}

	if handler := parent.PersistentPreRun; handler != nil {
		handler(cmd, args)
	}

	if handler := parent.PersistentPreRunE; handler != nil {
		if err := handler(cmd, args); err != nil {
			return err
		}
	}

	return nil
}
