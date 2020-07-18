package apispec

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

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
	AddHeaders([]ClientMeta) Client
	AddQueryParams([]ClientMeta) Client
	Exec(string) (*http.Response, error)
	ExecWithBody(string, io.ReadCloser) (*http.Response, error)
}

// Room struct
type Room struct {
	executed int
	client   Client
	reporter interface{}
}

// NewRoom creates an new test room
func NewRoom(client Client) *Room {
	return &Room{
		executed: 0,
		client:   client,
	}
}

// ExecuteTestSuite will made request and compare to expected responses
func (room *Room) ExecuteTestSuite(specs []SpecFile) error {
	// Loop throught all the spec files
	for _, spec := range specs {

		// Loop for every testing host
		for _, testing := range spec.Testing {

			url := testing.BuildURL()

			// Loop for every endpoint to test in the host
			for _, endpoint := range testing.Endpoints {
				var headers []string
				var queryParams []string
				var bodyRequest map[string]interface{}

				if endpoint.Request != nil {
					headers = endpoint.Request.Headers
					queryParams = endpoint.Request.QueryParams
					bodyRequest = endpoint.Request.Body["json"]
				}

				requestTo := endpoint.BuildPath(url)

				requestHeaders, err := fromColonSeparatedToClientMeta(headers)

				if err != nil {
					return err
				}

				requestQueryParams, err := fromColonSeparatedToClientMeta(queryParams)

				if err != nil {
					return err
				}

				marshaler := marshal.FactoryMarshal("json")
				body, err := marshaler.Marshal(bodyRequest)

				if err != nil {
					return err
				}

				response, err := room.client.
					AddURL(requestTo).
					AddHeaders(requestHeaders).
					AddQueryParams(requestQueryParams).
					ExecWithBody(endpoint.Method, body)

				if err != nil {
					return err
				}

				fmt.Println(response, err)
			}
		}
	}

	return nil
}

func fromColonSeparatedToClientMeta(colonSeparated []string) ([]ClientMeta, error) {
	totalColonSeparated := len(colonSeparated)

	if totalColonSeparated < 1 {
		return []ClientMeta{}, nil
	}

	meta := make([]ClientMeta, totalColonSeparated)

	for csIndex, cs := range colonSeparated {
		if separator := strings.Index(cs, ":"); separator == -1 {
			errmessage := fmt.Sprintf("Header %s must container header:value", cs)
			return []ClientMeta{}, errors.New(errmessage)
		}

		headerValues := strings.Split(cs, ":")

		meta[csIndex] = ClientMeta{
			Key:   headerValues[0],
			Value: headerValues[1],
		}
	}

	return meta, nil
}
