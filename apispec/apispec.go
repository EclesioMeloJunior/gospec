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

				requestTo := endpoint.BuildPath(url)

				headers, err := fromColonSeparatedToClientMeta(endpoint.Request.Headers)

				if err != nil {
					return err
				}

				queryParams, err := fromColonSeparatedToClientMeta(endpoint.Request.QueryParams)

				if err != nil {
					return err
				}

				values := endpoint.Request.Body["json"]

				marshaler := marshal.FactoryMarshal("json")
				body, err := marshaler.Marshal(values)

				if err != nil {
					return err
				}

				response, err := room.client.
					AddURL(requestTo).
					AddHeaders(headers).
					AddQueryParams(queryParams).
					ExecWithBody(http.MethodPost, body)

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
