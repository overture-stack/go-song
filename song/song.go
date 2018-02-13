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

package song

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"github.com/overture-stack/song-client/endpoint"
)

// Client struct allowing for making REST calls to a SONG server
type Client struct {
	accessToken string
	httpClient  *http.Client
	URLGenerator *endpoint.Endpoint
}

// CreateClient is a Factory Function for creating and returning a SONG client
func CreateClient(accessToken string, base *url.URL) *Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	httpClient := &http.Client{Transport: tr}
	songEndpoints := &endpoint.Endpoint{base}

	client := &Client{
		accessToken: accessToken,
		URLGenerator: songEndpoints,
		httpClient:  httpClient,
	}

	return client
}

// helper functions

func (c *Client) post(url url.URL, body []byte) string {
	var reader *bytes.Reader 

	if body == nil {
		reader = bytes.NewReader(body) 
	} else {
		reader = nil 
	}

        req, err := http.NewRequest("POST", url.String(), reader)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer " + c.accessToken)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return c.makeRequest(req)
}

func (c *Client) get(url url.URL) string {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	return c.makeRequest(req)
}

func (c *Client) makeRequest(req *http.Request) string {
	resp, err := c.httpClient.Do(req)
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
	return string(body)
}

// Upload uploads the file contents and returns the response
func (c *Client) Upload(studyID string, byteContent []byte, async bool) string {
        var url = c.URLGenerator.Upload(studyID, async)
        return c.post(url, byteContent) 
}

// GetStatus return the status JSON of an uploadID
func (c *Client) GetStatus(studyID string, uploadID string) string {
	var url = c.URLGenerator.GetStatus(studyID, uploadID)
        return c.get(url)
}

func (c *Client) GetServerStatus() string {
	var url = c.URLGenerator.IsAlive()
	return c.get(url)
}

// Save saves the specified uploadID as an analysis assuming it had passed validation
func (c *Client) Save(studyID string, uploadID string, ignoreCollisions bool) string {
	var url = c.URLGenerator.Save(studyID, uploadID, ignoreCollisions)
        return c.post(url,nil)
}

// Publish publishes a specified saved analysisID
func (c *Client) Publish(studyID string, analysisID string) string {
	var url = c.URLGenerator.Publish(studyID, analysisID)
        return c.post(url,nil)
}

func (c *Client) Suppress(studyID string, analysisID string) string {
	var url = c.URLGenerator.Suppress(studyID, analysisID)
	return c.post(url, nil)
}

func (c *Client) GetAnalysis(studyID string, analysisID string) string {
	var url = c.URLGenerator.GetAnalysis(studyID, analysisID)
	return c.post(url, nil)
}

func (c *Client) GetAnalysisFiles(studyID string, analysisID string) string {
	var url = c.URLGenerator.GetAnalysisFiles(studyID, analysisID)
	return c.get(url)
}

func (c *Client) IdSearch(studyID string, searchParams string) string {
	var url = c.URLGenerator.IdSearch(studyID, searchParams) 
	return c.get(url)
}

func (c *Client) InfoSearch(studyID string, includeInfo bool, searchTerms []string) string {
	var url = c.URLGenerator.InfoSearch(studyID, includeInfo, searchTerms)
	return c.get(url)
} 
