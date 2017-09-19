package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/andricdu/go-song/song"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(uploadCmd)
}

func upload(filePath string) {
	// init song client
	studyID, accessToken := viper.GetString("study"), viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}
	client := song.CreateClient(accessToken, songURL)

	// read the file
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
	}
	if len(b) < 1 {
		panic("File does not have any content!")
	}

	// use song client
	responseBody := client.Upload(studyID, b)
	fmt.Println(string(responseBody))
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload Analysis Metadata",
	Long:  `Uploads Metadata JSON describing an analysis and files for validation`,
	Run: func(cmd *cobra.Command, args []string) {
		upload(args[0])
	},
}
