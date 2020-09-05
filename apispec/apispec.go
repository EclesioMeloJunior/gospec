package apispec

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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
	GetURL() string

	AddHeaders(http.Header) Client
	GetHeaders() http.Header

	AddQueryParams([]ClientMeta) Client

	Exec(string) (*http.Response, error)
	ExecWithBody(string, io.ReadCloser) (*http.Response, error)
}

// Reporter interface
type Reporter interface {
	Results([]byte)
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
func (room *Room) ExecuteTestSuite(specs []SpecFile) ([][]byte, error) {
	results := [][]byte{}

	// Loop throught all the spec files
	for _, spec := range specs {
		room.executeSpecFileTesting(spec.Testing)
	}

	return results, nil
}

func (room *Room) executeSpecFileTesting(testing []SpecFileTesting) {
	for _, testing := range testing {
		fmt.Println(testing.Description)

		url := testing.BuildURL()

		room.executeTestingEndpoints(url, testing.Endpoints)
	}
}

func (room *Room) executeTestingEndpoints(url string, endpoints []TestingEndpoint) {
	for _, endpoint := range endpoints {
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
			log.Fatal(err)
			return
		}

		requestStartedAt := time.Now()

		fmt.Printf("%s - Start at %v\n", endpoint.Description, requestStartedAt)

		response, err := room.client.
			AddURL(requestTo).
			AddHeaders(headers).
			AddQueryParams(queryParams).
			ExecWithBody(endpoint.Method, body)

		if err != nil {
			log.Fatal(err)
			return
		}

		requestFinishedAt := time.Since(requestStartedAt)
		fmt.Printf("%s - Duration %v\n", endpoint.Description, requestFinishedAt)

		responseBody, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Fatal(err)
			return
		}

		defer response.Body.Close()

		expectedBody := endpoint.Expect.Body["json"]

		for headerKey, headerValue := range endpoint.Expect.GetHeaders() {
			responseHeader := response.Header[headerKey]

			for valueIndex, value := range responseHeader {
				if value != headerValue[valueIndex] {
					log.Println("Header does not match")
					log.Printf("From response: %s", value)
					log.Printf("Expected: %s\n", headerValue[valueIndex])
				}
			}
		}

		if expectBodies, ok := expectedBody.([]interface{}); ok {
			for expectBodyIndex, body := range expectBodies {
				fmt.Println(expectBodyIndex, body)
			}
		}

		fmt.Println(string(responseBody))
	}
}
