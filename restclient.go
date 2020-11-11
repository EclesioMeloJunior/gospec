package main

import (
	"net/http"
)

type httpclient interface {
	Do(*http.Request) (*http.Response, error)
}

type restclient interface {
	Get(string, http.Header) (*http.Response, error)
}

var (
	httpcli httpclient
)

func init() {
	httpcli = &http.Client{}
}

type client struct{}

func (c *client) Get(url string, headers http.Header) (*http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	request.Header = headers
	return httpcli.Do(request)
}
