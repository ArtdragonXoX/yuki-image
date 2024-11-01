package db

import (
	idb "yuki-image/internal/db"

	"github.com/spf13/cobra"
)

var DBReset = &cobra.Command{
	Use:   "reset",
	Short: "Reset the database to a clean state.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return idb.ResetDB()
	},
}

func init() {
	DBCmd.AddCommand(DBReset)
}
