package cmd

import (
	"yuki-image/internal/bootstrap"

	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the server",
	Long:  `Run the server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return bootstrap.Init()
	},
}

func init() {
	rootCmd.AddCommand(ServerCmd)
}
