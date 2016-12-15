package TelegramApiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BasicHttpClientWrapper struct {
	baseUrl      string
	postResponse func(originalResponse *http.Response, objectToReturn interface{}) (err error)
}

func (client *BasicHttpClientWrapper) Get(endpoint Endpoint, queryParams map[string]string, objectToReturn interface{}) error {
	queryValues := url.Values{}
	for k, v := range queryParams {
		queryValues.Add(k, v)
	}
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprint(client.baseUrl, endpoint), nil)
	req.URL.RawQuery = queryValues.Encode()
	return client.do(req, objectToReturn)
}

func (client *BasicHttpClientWrapper) Post(endpoint Endpoint, data interface{}, objectToReturn interface{}) error {
	requestBody, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprint(client.baseUrl, endpoint), bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	return client.do(req, objectToReturn)
}

func (client *BasicHttpClientWrapper) do(req *http.Request, objectToReturn interface{}) (err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = client.postResponse(resp, objectToReturn)
	return
}
