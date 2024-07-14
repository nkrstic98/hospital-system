package cmd

import (
	"hospital-system/server/config"
	"hospital-system/server/db"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var reinitCmd = &cobra.Command{
	Use:   "reinit",
	Short: "Recreates underlying database schema",
	Long:  "Recreates underlying database schema",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := db.OpenConnection(lo.FromPtrOr(cfg, config.DefaultConfig))
		if err != nil {
			panic(err)
		}

		err = db.ReinitDatabase()
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
	dbCmd.AddCommand(reinitCmd)
}
