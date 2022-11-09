package cmd

import (
	"mt-tgadmin/app"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs web GUI",

		RunE: func(cmd *cobra.Command, args []string) error {
			router := app.BuildWebRouter()

			http.ListenAndServe(
				":"+strconv.FormatUint(uint64(app.Global.Settings.GuiPort), 10),
				handlers.CompressHandler(router),
			)

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}
