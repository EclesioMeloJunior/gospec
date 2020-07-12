package main

import (
	"log"
	"os"

	"github.com/eclesiomelojunior/gospec/apispec"
	"github.com/eclesiomelojunior/gospec/config"
	"github.com/eclesiomelojunior/gospec/scanner"
)

func main() {
	conf := config.Load()
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

	err = apispec.ExecuteTestSuite(specfiles)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
		os.Exit(1)
	}
}
