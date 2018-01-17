package cmd

import (
	"fmt"
	"net/url"

	"github.com/andricdu/go-song/song"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(publishCmd)
}

func publish(analysisID string) {
	studyID, accessToken := viper.GetString("study"), viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}
	client := song.CreateClient(accessToken, songURL)
	responseBody := client.Publish(studyID, analysisID)
	fmt.Println(string(responseBody))
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a saved Analysis",
	Long:  `Publish a saved Analysis by specifying the AnalysisID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		publish(args[0])
	},
}
