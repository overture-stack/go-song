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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/url"
)
var searchTerms []string 
var includeInfo bool

func init() {
	RootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringArrayVarP(&searchTerms, "search-terms","t", []string{},"List of seach terms")
	searchCmd.Flags().StringP("analysis-id","a", "", "AnalysisID to match")
	searchCmd.Flags().StringP("donor-id", "d", "", "DonorID to match")
	searchCmd.Flags().StringP("file-id","f", "", "FileID to match")
	searchCmd.Flags().StringP("sample-id", "m", "", "SampleID to match")
	searchCmd.Flags().StringP("specimen-id", "p", "", "SampleID to match")
	searchCmd.Flags().BoolVarP(&includeInfo, "info", "n", false, "Include info field")
	
}

func search() {
	// init song client
	studyID := viper.GetString("study")
	client := createClient()

	// -a analysis-id, -d donor-id, -f file-id, -sa -sample-id, 
        // -sp specimen-id 
	// -t search-terms ([]), -i info(false)
        var responseBody string
	var params url.Values

	if  len(searchTerms) > 0 {
	    responseBody = client.InfoSearch(studyID, includeInfo, searchTerms)
        } else {
	    ids := [...]string{"sampleId", "specimenId", "donorId", "fileId"}
	    for _, id := range ids {
	   	val := viper.GetString(id)
                if val != "" {
		   params.Add(id,val)
                }
	    }
	    responseBody = client.IdSearch(studyID, params)
        }
	fmt.Println(string(responseBody))
}

var searchCmd = &cobra.Command{
	Use:   "search [options]",
	Short: "Suppress Analysis",
	Long:  `Suppresses an analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		search()
	},
}
