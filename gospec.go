package main

import (
	"fmt"
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
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println(contents)
}
