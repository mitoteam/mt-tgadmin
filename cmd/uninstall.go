package cmd

import (
	"log"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/mitoteam/mttools"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove installed system service " + app.Global.AppName,

		Run: func(cmd *cobra.Command, args []string) {
			if mttools.IsSystemdAvailable() {
				unitData := &mttools.ServiceData{
					Name: app.Global.Settings.ServiceName,
				}

				if err := unitData.UninstallSystemdService(); err != nil {
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

	rootCmd.AddCommand(cmd)
}
