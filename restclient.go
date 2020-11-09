package main

import (
	"net/http"
)

type httpclient interface {
	Do(*http.Request) (*http.Response, error)
}

var (
	client httpclient
)

func init() {
	client = &http.Client{}
}

func get(url string, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request.Header = headers
	return client.Do(request)
}
