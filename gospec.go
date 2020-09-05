package main

import (
	"log"
	"os"

	"github.com/eclesiomelojunior/gospec/apispec"
	"github.com/eclesiomelojunior/gospec/config"
	"github.com/eclesiomelojunior/gospec/httpclient"
	"github.com/eclesiomelojunior/gospec/scanner"
)

var (
	conf *config.Config
)

func init() {
	conf = config.Load()
}

func main() {
	sc := scanner.NewFileSystem(conf.ApispecFilesFlag)
	contents, err := apispec.LoadSpecFiles(sc)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}

	specfiles, err := apispec.ParseSpecFiles(contents)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}

	httpClient := httpclient.NewHTTPClient()

	testRoom := apispec.NewRoom(httpClient)

	_, err = testRoom.ExecuteTestSuite(specfiles)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}

}
