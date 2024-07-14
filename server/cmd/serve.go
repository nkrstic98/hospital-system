package cmd

import (
	"fmt"
	"hospital-system/server/app"
	"hospital-system/server/db"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts server",
	Long:  "Starts server",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := app.Build(*cfg)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := db.CloseConnection()
			if err != nil {
				panic(fmt.Errorf("failed to close database connection: %w", err))
			}

			cleanup()
		}()

		addr := fmt.Sprintf("%v:%v", cfg.Web.Host, cfg.Web.Port)
		slog.Info(fmt.Sprintf("Starting http server on port %v", cfg.Web.Port))
		go func() {
			if err = app.Run(addr); err != nil {
				panic(err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		slog.Info("Shutting down server...")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
