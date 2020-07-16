package apispec

import (
	"fmt"
	"net/http"
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
	AddURL(string) error
	AddHeaders([]ClientMeta) error
	AddQueryParams([]ClientMeta) error
	Exec(string) (*http.Response, error)
	ExecWithBody(string, map[interface{}]interface{}) (*http.Response, error)
}

// Room struct
type Room struct {
	executed int
	client   Client
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
	for _, spec := range specs {
		for _, testing := range spec.Testing {
			url := testing.BuildURL()

			for _, endpoint := range testing.Endpoints {
				requestTo := endpoint.BuildPath(url)
				headers := endpoint.Request.SetupHeaders()
				queryParams := endpoint.Request.SetupQueryParams()

				err := room.client.AddURL(requestTo)

				if err != nil {
					return err
				}

				err = room.client.AddHeaders(headers)

				if err != nil {
					return err
				}

				err = room.client.AddQueryParams(queryParams)

				if err != nil {
					return err
				}

				response, err := room.client.ExecWithBody(endpoint.Method, endpoint.Request.Body)

				if err != nil {
					return err
				}

				fmt.Println(response)
			}
		}
	}

	return nil
}
