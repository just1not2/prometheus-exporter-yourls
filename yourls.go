// Copyright: (c) 2022, Justin Béra (@just1not2) <me@just1not2.org>
// GNU General Public License v3.0+ (see LICENSE or https://www.gnu.org/licenses/gpl-3.0.txt)

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type YourlsClient struct {
	client    *http.Client
	url       string
	signature string
}

type StatBody struct {
	Stats   map[string]string `json:"stats"`
	Message string            `json:"message"`
	Code    float64           `json:"statusCode"`
}

func NewYourlsClient(url string, signature string, timeout time.Duration) *YourlsClient {
	return &YourlsClient{
		client:    &http.Client{Timeout: timeout},
		url:       url,
		signature: signature,
	}
}

func (client *YourlsClient) Request(parameters map[string]string) (*StatBody, error) {
	for parameter, value := range parameters {
		client.url += fmt.Sprintf("&%s=%s", parameter, value)
	}

	req, err := http.NewRequest("GET", client.url, nil)
	if err != nil {
		return nil, err
	}

	// Sends the HTTP request
	res, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var body StatBody

	// Sets the body into a JSON
	decoder := json.NewDecoder(res.Body)
	for {
		err := decoder.Decode(&body)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}

	if body.Code != 200 {
		return nil, fmt.Errorf("error %v: %s", body.Code, body.Message)
	}

	return &body, nil
}
