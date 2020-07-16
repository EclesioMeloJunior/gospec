package apispec

import (
	"fmt"
	"strings"
)

// EndpointRequest defines the endpoints request format
type EndpointRequest struct {
	Headers     []string                    `yaml:"headers"`
	QueryParams []string                    `yaml:"queryParams"`
	Body        map[interface{}]interface{} `yaml:"body"`
}

// EndpointExpect defines the expected from request
type EndpointExpect struct {
	Headers []string                    `yaml:"headers"`
	Status  int                         `yaml:"status"`
	Body    map[interface{}]interface{} `yaml:"body"`
}

// TestingEndpoint defines the endpoints to request
type TestingEndpoint struct {
	Path        string           `yaml:"path"`
	Description string           `yaml:"description"`
	Method      string           `yaml:"method"`
	Request     *EndpointRequest `yaml:"request"`
	Expect      *EndpointExpect  `yaml:"expect"`
}

// SpecFileTesting struct definition
type SpecFileTesting struct {
	Description     string            `yaml:"description,omitempty"`
	Protocol        string            `yaml:"protocol"`
	Port            int               `yaml:"port"`
	Host            string            `yaml:"host"`
	EndpointsPrefix string            `yaml:"endpointsPrefix,omitempty"`
	Endpoints       []TestingEndpoint `yaml:"endpoints"`
}

// SpecFile struct definition
type SpecFile struct {
	Version string            `yaml:"version"`
	Testing []SpecFileTesting `yaml:"testing"`
}

// IsValidVersion returns true if yaml has a valid version
func (sf SpecFile) IsValidVersion() bool {
	for _, a := range ValidVersions {
		if a == sf.Version {
			return true
		}
	}

	return false
}

// InvalidFields returns a list of invalid fields
func (sf *SpecFile) InvalidFields() []string {
	invalidFields := []string{}

	if sf.Version == "" {
		invalidFields = append(invalidFields, "Version is required")
	}

	if len(sf.Testing) < 1 {
		invalidFields = append(invalidFields, "Testing must have 1 or more items")
	}

	if len(sf.Testing) > 0 {
		for testingIndex, testing := range sf.Testing {
			if testing.Host == "" {
				invalidFieldMessage := fmt.Sprintf("Host at testing index %v is required", testingIndex)
				invalidFields = append(invalidFields, invalidFieldMessage)
			}

			if testing.Port < 1 {
				invalidFieldMessage := fmt.Sprintf("Port at testing index %v is required", testingIndex)
				invalidFields = append(invalidFields, invalidFieldMessage)
			}

			if testing.Protocol != "http" && testing.Protocol != "https" {
				invalidFieldMessage := fmt.Sprintf("Protocol at testing index %v must be http or https", testingIndex)
				invalidFields = append(invalidFields, invalidFieldMessage)
			}

			if len(testing.Endpoints) < 1 {
				invalidFieldMessage := fmt.Sprintf("Endpoints at testing index %v must have 1 or more items", testingIndex)
				invalidFields = append(invalidFields, invalidFieldMessage)
			}

			if len(testing.Endpoints) > 0 {
				for endpointIndex, endpoint := range testing.Endpoints {
					buildErrMessage := func(missingField string, message string) string {
						errmessage := fmt.Sprintf("%s at testing index %v, at endpoints index %v %s", missingField, testingIndex, endpointIndex, message)

						return errmessage
					}

					if endpoint.Path == "" {
						invalidFields = append(invalidFields, buildErrMessage("Path", "is required"))
					}

					if endpoint.Method == "" {
						invalidFields = append(invalidFields, buildErrMessage("Method", "is required"))
					}

					if !endpoint.IsValidMethod() {
						invalidFields = append(invalidFields, buildErrMessage("Method", fmt.Sprintf("allowed methods: %s", strings.Join(ValidMethods, ","))))
					}

					if endpoint.Request == nil {
						invalidFields = append(invalidFields, buildErrMessage("Request", "is required"))
					}

					if endpoint.Expect == nil {
						invalidFields = append(invalidFields, buildErrMessage("Expect", "is required"))
					}
				}
			}
		}
	}

	return invalidFields
}

// BuildURL returns a string with request url mounted
func (sft *SpecFileTesting) BuildURL() string {
	addPrefix := func(url string) string {
		prefix := sft.EndpointsPrefix

		if prefix == "" {
			return fmt.Sprintf("%s/", url)
		}

		if string(prefix[0]) == "/" {
			prefix = string(prefix[1:])
		}

		return fmt.Sprintf("%s/%s", url, prefix)
	}

	url := fmt.Sprintf("%s://%s:%v", sft.Protocol, sft.Host, sft.Port)
	return addPrefix(url)
}

// IsValidMethod is to validate http method
func (endpoint *TestingEndpoint) IsValidMethod() bool {
	for _, method := range ValidMethods {
		if method == endpoint.Method {
			return true
		}
	}

	return false
}

// BuildPath will mount the complete URL with endpoint to request
func (endpoint *TestingEndpoint) BuildPath(url string) string {
	path := endpoint.Path

	if string(path[0]) == "/" {
		path = path[1:]
	}

	return fmt.Sprintf("%s%s", url, path)
}

// SetupHeaders will normalize headers
func (request *EndpointRequest) SetupHeaders() []ClientMeta {
	return nil
}

// SetupQueryParams will normalize queryparams
func (request *EndpointRequest) SetupQueryParams() []ClientMeta {
	return nil
}
