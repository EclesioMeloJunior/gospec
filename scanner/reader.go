package scanner

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eclesiomelojunior/gospec/apispec"
)

// FileScanner struct definition
type FileScanner struct {
	filepath string
}

// New returns an pointer to FileScanner struct
func New(filepath string) *FileScanner {
	return &FileScanner{
		filepath: filepath,
	}
}

// NewFileSystem returns the apispec.FileSystem interface
func NewFileSystem(filepath string) apispec.FileSystem {
	return New(filepath)
}

// IsDir returns if filepath is a directory
func (fs FileScanner) IsDir(filepath string) bool {
	info, err := os.Stat(filepath)

	if err != nil {
		log.Fatalln(err.Error())
		return false
	}

	return info.IsDir()
}

// GetFiles return all files inside filepath
func (fs FileScanner) GetFiles() []string {
	files := make([]string, 0)

	err := filepath.Walk(fs.filepath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err.Error())
		return []string{}
	}

	return files
}

// GetLocation return the filepath
func (fs FileScanner) GetLocation() string {
	return fs.filepath
}

// IsYamlFile return if a filepath is a yaml
func (fs FileScanner) IsYamlFile(filepathname string) bool {
	fileinfo, err := os.Stat(filepathname)

	if err != nil {
		log.Fatalln(err.Error())
		return false
	}

	if fileinfo.IsDir() {
		return false
	}

	extension := filepath.Ext(filepathname)

	if extension == ".yaml" || extension == ".yml" {
		return true
	}

	return false
}

// ReadFile returns the content of a filepath or error
func (fs FileScanner) ReadFile(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}
