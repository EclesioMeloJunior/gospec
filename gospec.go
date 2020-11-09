package main

import "errors"

type testendpoints struct {
	Path   string
	Method string
}

type testsection struct {
	Port      int
	Host      string
	Protocol  string
	Endpoints []testendpoints
}

type specfile struct {
	Version int
	Testing []testsection
}

func (s *specfile) isVersionValid() bool {
	return s.Version == 1
}

func gospec(testfile string) error {
	if testfile == "" {
		return errors.New("Test file could not be empty")
	}

	return nil
}

func parseFile(testfile string) *specfile {
	return &specfile{}
}
