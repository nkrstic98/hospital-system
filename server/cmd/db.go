package cmd

import (
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Runs command on the database resource",
	Long:  "Runs command on the database resource",
}

func init() {
	rootCmd.AddCommand(dbCmd)
}
