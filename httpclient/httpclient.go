package httpclient

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/eclesiomelojunior/gospec/apispec"
)

type client struct {
	url     string
	method  string
	headers map[string]string
}

// NewHTTPClient return the interface apispec.Client
func NewHTTPClient() apispec.Client {
	return &client{}
}

func (c *client) AddURL(url string) error {
	if url == "" {
		errmessage := fmt.Sprintf("URL cannot be empty")
		return errors.New(errmessage)
	}

	c.url = url

	return nil
}

func (c *client) AddHeaders(headers []apispec.ClientMeta) error {
	return nil

}

func (c *client) AddQueryParams(qp []apispec.ClientMeta) error {
	return nil
}

func (c *client) Exec(method string) (*http.Response, error) {
	return nil, nil
}

func (c *client) ExecWithBody(method string, body map[interface{}]interface{}) (*http.Response, error) {
	return nil, nil
}
