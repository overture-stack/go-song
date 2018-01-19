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
	"path"
	"time"
)

// Client struct allowing for making REST calls to a SONG server
type Client struct {
	accessToken string
	songURL     *url.URL
	httpClient  *http.Client
}

// CreateClient is a Factory Function for creating and returning a SONG client
func CreateClient(accessToken string, songURL *url.URL) *Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	httpClient := &http.Client{Transport: tr}

	client := &Client{
		accessToken: accessToken,
		songURL:     songURL,
		httpClient:  httpClient,
	}

	return client
}

// Upload uploads the file contents and returns the response
func (c *Client) Upload(studyID string, byteContent []byte) string {
	requestURL := *c.songURL
	requestURL.Path = path.Join(c.songURL.Path, "upload", studyID)
	req, err := http.NewRequest("POST", requestURL.String(), bytes.NewReader(byteContent))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Content-Type", "application/json")
	return c.makeRequest(req)
}

// GetStatus return the status JSON of an uploadID
func (c *Client) GetStatus(studyID string, uploadID string) string {
	requestURL := *c.songURL
	requestURL.Path = path.Join(c.songURL.Path, "upload", studyID, "status", uploadID)
	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	return c.makeRequest(req)
}

// Save saves the specified uploadID as an analysis assuming it had passed validation
func (c *Client) Save(studyID string, uploadID string) string {
	requestURL := *c.songURL
	requestURL.Path = path.Join(c.songURL.Path, "upload", studyID, "save", uploadID)
	req, err := http.NewRequest("POST", requestURL.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	return c.makeRequest(req)
}

// Publish publishes a specified saved analysisID
func (c *Client) Publish(studyID string, analysisID string) string {
	requestURL := *c.songURL
	requestURL.Path = path.Join(c.songURL.Path, "studies", studyID, "publish", analysisID)
	req, err := http.NewRequest("POST", requestURL.String(), nil)
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
