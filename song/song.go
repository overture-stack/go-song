package song

import (
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

// GetStatus return the status JSON of an uploadID
func (c *Client) GetStatus(studyID string, uploadID string) string {

	requestURL := *c.songURL

	requestURL.Path = path.Join(c.songURL.Path, "upload", studyID, "status", uploadID)
	req, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
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
