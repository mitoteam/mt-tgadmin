package cmd

import (
	"log"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/mitoteam/mt-tgadmin/mttools"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Creates system service to run " + app.Global.AppName,

		Run: func(cmd *cobra.Command, args []string) {
			if mttools.IsSystemdAvailable() {
				unitData := &mttools.ServiceData{
					Name:      app.Global.Settings.ServiceName,
					User:      app.Global.Settings.ServiceUser,
					Group:     app.Global.Settings.ServiceGroup,
					Autostart: app.Global.ServiceAutostart,
				}

				if err := unitData.InstallSystemdService(); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Fatalf(
					"Directory %s does not exists. Only systemd based services supported for now.\n",
					mttools.SystemdServiceDirPath,
				)
			}
		},
	}

	cmd.PersistentFlags().BoolVar(
		&app.Global.ServiceAutostart,
		"autostart",
		false,
		"Set service to be auto started after boot. Please note this does not auto starts service after installation.",
	)

	rootCmd.AddCommand(cmd)
}
