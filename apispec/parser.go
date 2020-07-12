package apispec

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

// SpecFileTesting struct definition
type SpecFileTesting struct {
	Description     string `yaml:"description,omitempty"`
	Protocol        string `yaml:"protocol"`
	Port            int    `yaml:"port"`
	Host            string `yaml:"host"`
	EndpointsPrefix string `yaml:"endpointsPrefix,omitempty"`
}

// SpecFile struct definition
type SpecFile struct {
	Version string            `yaml:"version"`
	Testing []SpecFileTesting `yaml:"testing"`
}

// ParseSpecFiles transform the yaml content int defined structs
func ParseSpecFiles(fileContents [][]byte) ([]SpecFile, error) {
	if len(fileContents) < 1 {
		return []SpecFile{}, nil
	}

	specfiles := make([]SpecFile, len(fileContents))

	for specfileIndex, fileContent := range fileContents {
		specfile := &SpecFile{}
		err := yaml.Unmarshal(fileContent, specfile)

		if err != nil {
			return nil, err
		}

		if !specfile.IsValidVersion() {
			errmessage := fmt.Sprintf("Spec file has an invalid version: %s", specfile.Version)
			return nil, errors.New(errmessage)
		}

		invalidFields := specfile.InvalidFields()

		if len(invalidFields) > 0 {
			errmessage := strings.Join(invalidFields, "\n")
			return nil, errors.New(errmessage)
		}

		specfiles[specfileIndex] = *specfile
	}

	return specfiles, nil
}

// IsValidVersion returns true if yaml has a valid version
func (sf *SpecFile) IsValidVersion() bool {
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
		}
	}

	return invalidFields
}
