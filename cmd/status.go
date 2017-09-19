package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RootCmd.AddCommand(statusCmd)
}

func createClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	client := &http.Client{Transport: tr}
	return client
}

func getStatus(uploadID string) {

	client := createClient()

	study, accessToken := viper.GetString("study"), viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}

	songURL.Path = path.Join(songURL.Path, "upload", study, "status", uploadID)
	url := songURL.String()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// To compare status codes, you should always use the status constants
	// provided by the http package.
	if resp.StatusCode != http.StatusOK {
		panic("Request was not OK: " + resp.Status)
	}

	// Example of JSON decoding on a reader.
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of uploaded analysis",
	Long:  `Get status of uploaded analysis`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getStatus(args[0])
	},
}
