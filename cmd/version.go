package cmd

import (
	"fmt"

	"github.com/mitoteam/mt-tgadmin/app"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Prints the raw version number of " + app.Global.AppName + ".",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(app.Global.Version)
		},
	})
}
