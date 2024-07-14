package cmd

import (
	"hospital-system/server/config"
	"hospital-system/server/db"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Populate database with predefined data",
	Long:  "Populate database with predefined data",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := db.OpenConnection(lo.FromPtrOr(cfg, config.DefaultConfig))
		if err != nil {
			panic(err)
		}

		err = db.SeedDatabase()
		if err != nil {
			panic(err)
		}

		err = db.CloseConnection()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	dbCmd.AddCommand(seedCmd)
}
