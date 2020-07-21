package apispec

import (
	"io"
	"net/http"

	"github.com/eclesiomelojunior/gospec/marshal"
)

// ValidVersions has all supported versions
var ValidVersions = []string{
	"1.0",
}

// ValidMethods has all supported versions
var ValidMethods = []string{
	"GET", "POST", "PUT", "DELETE", "OPTIONS",
}

// ClientMeta standartize headers
type ClientMeta struct {
	Key   string
	Value string
}

// Client interface standartize the clients
type Client interface {
	AddURL(string) Client
	AddHeaders(http.Header) Client
	AddQueryParams([]ClientMeta) Client
	Exec(string) (*http.Response, error)
	ExecWithBody(string, io.ReadCloser) (*http.Response, error)
}

// Reporter interface
type Reporter interface {
	Results([]byte)
}

// Assert interface
type Assert interface {
	TestBody(interface{}) Assert
	TestHeaders(http.Header) Assert
	TestStatusCode(int) Assert

	Test(*http.Response) ([]byte, error)
}

// Room struct
type Room struct {
	executed int
	client   Client
	assert   Assert
}

// NewRoom creates an new test room
func NewRoom(client Client, assert Assert) *Room {
	return &Room{
		executed: 0,
		client:   client,
		assert:   assert,
	}
}

// ExecuteTestSuite will made request and compare to expected responses
func (room *Room) ExecuteTestSuite(specs []SpecFile) ([][]byte, error) {
	results := [][]byte{}

	// Loop throught all the spec files
	for _, spec := range specs {

		// Loop for every testing host
		for _, testing := range spec.Testing {

			url := testing.BuildURL()

			// Loop for every endpoint to test in the host
			for _, endpoint := range testing.Endpoints {

				var (
					headers     http.Header
					queryParams []ClientMeta
					bodyRequest map[string]interface{}
				)

				if endpoint.Request != nil {
					headers = endpoint.Request.GetHeaders()
					queryParams = endpoint.Request.GetQueryParams()
					bodyRequest = endpoint.Request.Body["json"]
				}

				requestTo := endpoint.BuildPath(url)

				marshaler := marshal.FactoryMarshal("json")
				body, err := marshaler.Marshal(bodyRequest)

				if err != nil {
					return nil, err
				}

				response, err := room.client.
					AddURL(requestTo).
					AddHeaders(headers).
					AddQueryParams(queryParams).
					ExecWithBody(endpoint.Method, body)

				if err != nil {
					return nil, err
				}

				result, err := room.assert.
					TestBody(endpoint.Expect.Body).
					TestHeaders(endpoint.Expect.GetHeaders()).
					TestStatusCode(endpoint.Expect.Status).
					Test(response)

				if err != nil {
					return nil, err
				}

				results = append(results, result)
			}
		}
	}

	return results, nil
}
