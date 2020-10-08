package cmd

import (
	"fmt"
	"github.com/bamItsCam/cardtage/internal/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cardtage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("v%s.%s.%s", version.Major, version.Minor, version.Patch)
	},
}
