package cmd

import (
	"github.com/spf13/cobra"
	"hospital-system/authorization/app"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts server",
	Long:  "Starts server",
	Run: func(cmd *cobra.Command, args []string) {
		app, _, err := app.Build(*cfg)
		if err != nil {
			panic(err)
		}

		//addr := fmt.Sprintf("%v:%v", cfg.Web.Host, cfg.Web.Port)
		//slog.Info(fmt.Sprintf("Starting http server on port %v", cfg.Web.Port))
		//go func() {
		//	if err = app.Run(addr); err != nil {
		//		panic(err)
		//	}
		//}()
		//
		//quit := make(chan os.Signal, 1)
		//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		//<-quit
		//
		//slog.Info("Shutting down server...")

		if err := app.Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
