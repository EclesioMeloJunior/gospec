package main

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var MockedDo func(*http.Request) (*http.Response, error)

type httpclientMock struct{}

func (h *httpclientMock) Do(r *http.Request) (*http.Response, error) {
	return MockedDo(r)
}

func mockHTTPpClient(doimpl func(*http.Request) (*http.Response, error)) httpclient {
	m := &httpclientMock{}
	MockedDo = doimpl

	return m
}

func TestGetWithSuccesfullResponse(t *testing.T) {
	httpcli = mockHTTPpClient(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, r.Method, http.MethodGet)
		assert.Nil(t, r.Header)

		resp := &http.Response{}
		resp.StatusCode = http.StatusOK

		return resp, nil
	})

	var c restclient
	c = &client{}

	resp, err := c.Get("http://someurl.com", nil)

	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestGetWhenHttpClientDoFails(t *testing.T) {
	httpcli = mockHTTPpClient(func(r *http.Request) (*http.Response, error) {
		return nil, errors.New("Problems to execute the request")
	})

	c := &client{}
	resp, err := c.Get("http://some.com", nil)

	assert.Nil(t, resp)
	assert.Equal(t, err.Error(), "Problems to execute the request")
}
