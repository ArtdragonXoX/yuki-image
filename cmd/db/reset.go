package db

import (
	idb "yuki-image/internal/db"

	"github.com/spf13/cobra"
)

var DBReset = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database to a clean state.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return idb.InitDataBase()
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		return idb.ResetTable()
	},
}

func init() {
	DBCmd.AddCommand(DBReset)
}
