package apispec

import (
	"errors"
	"fmt"
)

// FileSystem defines where to search
type FileSystem interface {
	IsDir(filepath string) bool
	GetFiles() []string
	GetLocation() string
	IsYamlFile(filepath string) bool
	ReadFile(filepath string) ([]byte, error)
}

// LoadSpecFiles will read all yaml files
func LoadSpecFiles(apispecFS FileSystem) ([][]byte, error) {

	if !apispecFS.IsDir(apispecFS.GetLocation()) {
		return nil, errors.New("Cannot load spec files")
	}

	files := apispecFS.GetFiles()

	if len(files) < 1 {
		errmessage := fmt.Sprintf("There is 0 files at %s", apispecFS.GetLocation())
		return nil, errors.New(errmessage)
	}

	allReadFiles := make([][]byte, len(files))

	for fileIndex, file := range files {
		if !apispecFS.IsYamlFile(file) || apispecFS.IsDir(file) {
			errmessage := fmt.Sprintf("The file %s is not in yaml format", file)
			return nil, errors.New(errmessage)
		}

		content, err := apispecFS.ReadFile(file)

		if err != nil {
			return nil, err
		}

		allReadFiles[fileIndex] = content
	}

	return allReadFiles, nil
}
