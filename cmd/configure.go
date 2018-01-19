/*
 *     Copyright (C) 2018  Ontario Institute for Cancer Research
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU Affero General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU Affero General Public License for more details.
 *
 *     You should have received a copy of the GNU Affero General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package cmd

import (
	"fmt"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(configureCmd)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func verifyPath(fullPath string) {
	_, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		fmt.Println("No configuration existing configuration file, creating new config.")
	} else {
		fmt.Println("Existing configuration found. Type 'y' to continue...")
		var input string
		fmt.Scanln(&input)
		if input != "y" {
			os.Exit(0)
		}
	}
}

func doConfigure() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fullPath := path.Join(home, ".song.yaml")
	verifyPath(fullPath)

	file, err := os.Create(fullPath)
	check(err)
	defer file.Close()

	var accessToken string
	fmt.Println("Please enter your access token: ")
	fmt.Scanln(&accessToken)

	var songURL string
	fmt.Println("Please enter URL of SONG server: ")
	fmt.Scanln(&songURL)

	var study string
	fmt.Println("Please enter study ID: ")
	fmt.Scanln(&study)

	accessTokenConfig := "accessToken: " + accessToken + "\n"
	songURLConfig := "songURL: " + songURL + "\n"
	studyConfig := "study: " + study + "\n"

	_, err = file.WriteString(accessTokenConfig)
	check(err)
	_, err = file.WriteString(songURLConfig)
	check(err)
	_, err = file.WriteString(studyConfig)
	check(err)
	file.Sync()
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "configures SONG client",
	Long:  `Sets configuration values in config file.`,
	Run: func(cmd *cobra.Command, args []string) {
		doConfigure()
	},
}
