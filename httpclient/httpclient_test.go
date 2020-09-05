package httpclient_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eclesiomelojunior/gospec/apispec"
	"github.com/eclesiomelojunior/gospec/httpclient"
	"github.com/eclesiomelojunior/gospec/tests"
)

func TestHTTPClientAddURL(t *testing.T) {
	client := httpclient.NewHTTPClient()
	expectedURL := "anyurl"

	client.AddURL(expectedURL)

	if client.GetURL() != expectedURL {
		tests.ExpectationError(t, expectedURL, client.GetURL(), "client.GetURL()")
	}
}

func TestHTTPClientAddHeaders(t *testing.T) {
	client := httpclient.NewHTTPClient()
	headers := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {"mytoken"},
	}

	client.AddHeaders(headers)

	if len(client.GetHeaders()) == 0 {
		tests.ExpectationError(
			t,
			headers,
			client.GetHeaders(),
			fmt.Sprintf("client.GetHeaders()"),
		)
	}

	for headerName, values := range client.GetHeaders() {
		expectedvalues := headers[headerName]

		if !reflect.DeepEqual(expectedvalues, values) {
			tests.ExpectationError(
				t,
				expectedvalues,
				values,
				fmt.Sprintf("client.GetHeaders()"),
			)
		}
	}
}

func TestHTTPClientAddQueryParamsWithNoURL(t *testing.T) {
	client := httpclient.NewHTTPClient()

	expected := "/?query=param&another=param"
	client.AddQueryParams([]apispec.ClientMeta{
		{
			Key:   "query",
			Value: "param",
		},
		{
			Key:   "another",
			Value: "param",
		},
	})

	if client.GetURL() != expected {
		tests.ExpectationError(
			t,
			expected,
			client.GetURL(),
			"client.AddQueryParams()",
		)
	}
}

func TestHTTPClientAddQueryParamsWithURL(t *testing.T) {
	client := httpclient.NewHTTPClient()
	expected := "http://url.com?another=param"

	client.AddURL("http://url.com")
	client.AddQueryParams([]apispec.ClientMeta{
		{
			Key:   "another",
			Value: "param",
		},
	})

	if client.GetURL() != expected {
		tests.ExpectationError(
			t,
			expected,
			client.GetURL(),
			"client.AddQueryParams()",
		)
	}
}

func TestHTTPClientAddQueryParamsWithNoParams(t *testing.T) {
	client := httpclient.NewHTTPClient()
	expected := "http://url.com"

	client.AddURL("http://url.com")
	client.AddQueryParams([]apispec.ClientMeta{})

	if client.GetURL() != expected {
		tests.ExpectationError(
			t,
			expected,
			client.GetURL(),
			"client.AddQueryParams()",
		)
	}
}
