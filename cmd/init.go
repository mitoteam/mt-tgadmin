package cmd

import (
	"errors"
	"fmt"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/mitoteam/mttools"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Creates bot settings file with defaults in working directory. --settings option can be used to specify settings file name or location explicitly.",

		RunE: func(cmd *cobra.Command, args []string) error {
			filename := app.Global.SettingsFilename

			if mttools.IsFileExists(filename) {
				return errors.New("can not initialize existing file: " + filename)
			}

			comment := `
File created automatically by '` + app.Global.AppName + ` init' command. There are all available
options listed here with its default values. Recommendation is to edit options you
want to change and remove all others to keep this as simple as possible.
`

			if err := mttools.SaveYamlSettingToFile(filename, comment, &app.Global.Settings); err != nil {
				return err
			}

			fmt.Println("Default settings written to " + filename)

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}
