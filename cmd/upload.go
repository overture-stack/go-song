package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload Analysis Metadata",
	Long:  `Uploads Metadata JSON describing an analysis and files for validation`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1 Alpha")
	},
}
