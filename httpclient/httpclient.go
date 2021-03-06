package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/eclesiomelojunior/gospec/apispec"
)

type client struct {
	url        string
	method     string
	headers    map[string][]string
	httpclient *http.Client
}

// NewHTTPClient return the interface apispec.Client
func NewHTTPClient() apispec.Client {
	return &client{
		httpclient: &http.Client{},
	}
}

func (c *client) AddURL(url string) apispec.Client {
	c.url = url
	return c
}

func (c *client) GetURL() string {
	return c.url
}

func (c *client) AddHeaders(headers http.Header) apispec.Client {
	c.headers = headers
	return c
}

func (c *client) GetHeaders() http.Header {
	return c.headers
}

func (c *client) AddQueryParams(queryParams []apispec.ClientMeta) apispec.Client {
	if c.url == "" {
		c.url = "/"
	}

	if len(queryParams) < 1 {

		if qpIndex := strings.Index(c.url, "?"); qpIndex != -1 {
			c.url = string(c.url[:qpIndex])
		}

		return c
	}

	formatedQueryParams := make([]string, len(queryParams))

	for queryParamIndex, queryParam := range queryParams {
		formatedQueryParams[queryParamIndex] = fmt.Sprintf("%s=%s", queryParam.Key, queryParam.Value)
	}

	urlQueryParams := strings.Join(formatedQueryParams, "&")

	if len(c.url) > 1 && string(c.url[len(c.url)-1]) == "/" {
		c.url = fmt.Sprintf("%s?%s", string(c.url[:len(c.url)-1]), urlQueryParams)
		return c
	}

	c.url = fmt.Sprintf("%s?%s", string(c.url), urlQueryParams)
	return c
}

func (c *client) Exec(method string) (*http.Response, error) {
	return nil, nil
}

func (c *client) ExecWithBody(method string, body io.ReadCloser) (*http.Response, error) {
	URL, err := url.Parse(c.url)

	if err != nil {
		return nil, err
	}

	request := &http.Request{
		Method: method,
		Header: c.headers,
		URL:    URL,
		Body:   body,
	}

	return c.httpclient.Do(request)
}
