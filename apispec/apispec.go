package apispec

import (
	"net/http"
)

// ValidVersions has all supported versions
var ValidVersions = []string{
	"1.0",
}

// ClientMeta standartize headers
type ClientMeta struct {
	Key   string
	Value string
}

// Client interface standartize the clients
type Client interface {
	AddHeaders([]ClientMeta)
	AddQueryParam([]ClientMeta)
	Exec() (*http.Response, error)
}

// ExecuteTestSuite will made request and compare to expected responses
func ExecuteTestSuite(specs []SpecFile) error {
	for _, spec := range specs {
		for _, testing := range spec.Testing {
			url := testing.BuildURL()

			for _, endpoint := range testing.Endpoints {
				requestTo := endpoint.BuildPath(url)
				headers := endpoint.Request.SetupHeaders()
				queryParams := endpoint.Request.SetupQueryParams()
			}
		}
	}

	return nil
}
