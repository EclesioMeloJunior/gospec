package apispec

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

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
