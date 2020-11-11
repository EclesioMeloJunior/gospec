package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	testFile string
)

func init() {
	flag.StringVar(&testFile, "file", "", "set the test file to use")
}

func main() {
	flag.Parse()

	if err := gospec(testFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
