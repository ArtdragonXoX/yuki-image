package cmd

import (
	"fmt"
	"os"
	"yuki-image/cmd/db"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yuki",
	Short: "yuki-image is a simple local gallery service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run yuki-image")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(db.DBCmd)
}
