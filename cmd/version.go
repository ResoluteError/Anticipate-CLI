package cmd

import (
	"anticipate/env"
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Anticipate",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(env.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
