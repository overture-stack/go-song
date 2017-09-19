package cmd

import (
	"fmt"
	"net/url"

	"github.com/andricdu/go-song/song"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(saveCmd)
}

func save(uploadID string) {
	studyID, accessToken := viper.GetString("study"), viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}
	client := song.CreateClient(accessToken, songURL)
	responseBody := client.Save(studyID, uploadID)
	fmt.Println(string(responseBody))
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save the uploaded Analysis",
	Long:  `Save the uploaded Analysis`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		save(args[0])
	},
}
