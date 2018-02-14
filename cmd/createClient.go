package cmd

import (
	"net/url"
	"github.com/overture-stack/song-client/song"
	"github.com/spf13/viper"
)
func createClient() *song.Client {
	accessToken := viper.GetString("accessToken")
	songURL, err := url.Parse(viper.GetString("songURL"))
	if err != nil {
		panic(err)
	}
	client := song.CreateClient(accessToken, songURL)
	return client
}
