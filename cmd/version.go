package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version number",
	Long:  `Prints the version of this SONG CLI Application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1 Alpha")
	},
}
