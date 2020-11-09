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
	flag.StringVar(&testFile, "test", "", "set the test file to use")
	flag.Parse()
}

func main() {
	if err := gospec(testFile); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
