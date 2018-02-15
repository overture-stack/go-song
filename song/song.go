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
	//"fmt"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
)

// Client struct allowing for making REST calls to a SONG server
type Client struct {
	accessToken string
	httpClient  *http.Client
	endpoint *Endpoint
}

type EndpointHandler func(... string) *http.Request 

// CreateClient is a Factory Function for creating and returning a SONG client
func CreateClient(accessToken string, base *url.URL) *Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	httpClient := &http.Client{Transport: tr}
	songEndpoints := &Endpoint{base}

	client := &Client{
		accessToken: accessToken,
		endpoint: songEndpoints,
		httpClient:  httpClient,
	}

	return client
}

func (c *Client) post(address url.URL, body []byte) string {
        req  := createRequest("POST", address, body)

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return c.makeRequest(req)
}

func (c *Client) put(address url.URL, body []byte) string {
	req := createRequest("PUT", address, body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return c.makeRequest(req)
}

func (c *Client) get(address url.URL) string {
	req := createRequest("GET", address, nil)
	return c.makeRequest(req)
}

func createRequest(requestType string, address url.URL, body []byte) *http.Request {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(requestType, address.String(), nil)
	} else {
		req, err = http.NewRequest(requestType, address.String(), bytes.NewReader(body))
	}

	if err != nil {
		panic(err)
	}

	return req
}

func (c *Client) makeRequest(req *http.Request) string {
	req.Header.Add("Authorization", "Bearer " + c.accessToken)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		panic("Request was not OK: " + resp.Status + string(body)) 
	}

	return string(body)
}

// Upload uploads the file contents and returns the response
func (c *Client) Upload(studyID string, byteContent []byte, async bool) string {
        var url = c.endpoint.Upload(studyID, async)
        return c.post(url, byteContent) 
}

// GetStatus return the status JSON of an uploadID
func (c *Client) GetStatus(studyID string, uploadID string) string {
	return c.get(c.endpoint.GetStatus(studyID, uploadID))
}

func (c *Client) GetServerStatus() string {
	return c.get(c.endpoint.IsAlive())
}

// Save saves the specified uploadID as an analysis assuming it had passed validation
func (c *Client) Save(studyID string, uploadID string, ignoreCollisions bool) string {
	return c.post(c.endpoint.Save(studyID, uploadID, ignoreCollisions), nil)
}

// Publish publishes a specified saved analysisID
func (c *Client) Publish(studyID string, analysisID string) string {
	return c.put(c.endpoint.Publish(studyID, analysisID),nil)
}

func (c *Client) Suppress(studyID string, analysisID string) string {
	return c.put(c.endpoint.Suppress(studyID, analysisID),nil)
}

func (c *Client) getAnalysis(studyID string, analysisID string) string {
	return c.get(c.endpoint.GetAnalysis(studyID, analysisID))
}

func (c *Client) getAnalysisFiles(studyID string, analysisID string) string {
	return c.get(c.endpoint.GetAnalysisFiles(studyID, analysisID))
}

func (c *Client) IdSearch(studyID string, ids map[string]string) string {
	searchTerms, err := json.Marshal(ids)
	if err != nil {
		panic(err)
	}
	return c.post(c.endpoint.IdSearch(studyID), searchTerms)
}

type InfoKey struct {
	Key string  `json:"key"`
	Value string `json:"value"`
}

type InfoSearchRequest struct {
	IncludeInfo bool `json:"includeInfo"`
	SearchTerms []InfoKey `json:"searchTerms"`
}

func createInfoSearchRequest(includeInfo bool, terms map[string]string) InfoSearchRequest {
	var searchTerms = []InfoKey{}
	for k,v := range terms {
		searchTerms=append(searchTerms, InfoKey{k,v})
	}
	return InfoSearchRequest{includeInfo, searchTerms}
}

func (c *Client) InfoSearch(studyID string, includeInfo bool, terms map[string]string) string {
	data := createInfoSearchRequest(includeInfo, terms)
	searchRequest, err := json.Marshal(data) 

	if err != nil {
		panic(err)
	}
	return c.post(c.endpoint.InfoSearch(studyID), searchRequest)
}


func (c *Client) Manifest(studyID string, analysisID string) string {
	var data = c.getAnalysisFiles(studyID,analysisID)	
	manifest := createManifest(analysisID, data)
	return manifest
}
