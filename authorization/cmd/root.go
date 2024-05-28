package cmd

import (
	"github.com/spf13/cobra"
	"hospital-system/authorization/config"
	"os"
)

var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "authorization",
	Short: "Authorization microservice",
	Long:  "Authorization microservice",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfg = &config.DefaultConfig
}
