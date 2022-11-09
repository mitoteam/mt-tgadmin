package cmd

import (
	"context"
	"log"
	"mt-tgadmin/app"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Runs web GUI",

		RunE: func(cmd *cobra.Command, args []string) error {
			router := app.BuildWebRouter()

			//Graceful shutdown according to https://github.com/gorilla/mux#graceful-shutdown
			httpSrv := &http.Server{
				Addr:         ":" + strconv.FormatUint(uint64(app.Global.Settings.GuiPort), 10),
				WriteTimeout: time.Second * 10,
				ReadTimeout:  time.Second * 20,
				IdleTimeout:  time.Second * 60,
				Handler:      router,
			}

			log.Println("Starting up web server. Press Ctrl + C to stop it.")

			go func() {
				if err := httpSrv.ListenAndServe(); err != nil {
					log.Fatalln(err)
				}
			}()

			c := make(chan os.Signal, 1)

			// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
			// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
			signal.Notify(c, os.Interrupt)

			// Block until we receive our signal.
			<-c

			// Create a deadline to wait for (10s).
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			// Doesn't block if no connections, but will otherwise wait
			// until the timeout deadline.
			httpSrv.Shutdown(ctx)

			log.Println("Shutting down web server")

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}
