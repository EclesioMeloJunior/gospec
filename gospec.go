package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

// testoutput standartize [10-10-2020] Test http://localhost:8000/endpoint: OK
const testoutput = `[%s] Test %s: %s`

type testexpected struct {
	Headers []string               `yaml:"headers"`
	Status  int                    `yaml:"status"`
	Body    map[string]interface{} `yaml:"body"`
}

type testendpoints struct {
	Path   string `yaml:"path"`
	Method string `yaml:"method"`
}

type testsection struct {
	Port      int             `yaml:"port"`
	Host      string          `yaml:"host"`
	Protocol  string          `yaml:"protocol"`
	Endpoints []testendpoints `yaml:"endpoints"`
}

type specfile struct {
	Version int           `yaml:"version"`
	Testing []testsection `yaml:"testing"`
}

func (s *specfile) isVersionValid() bool {
	return s.Version == 1
}

func gospec(testfile string) error {
	if testfile == "" {
		return errors.New("Test file could not be empty")
	}

	spec, err := parseFile(testfile)

	if err != nil {
		return err
	}

	suite(spec)

	return nil
}

func parseFile(testfile string) (*specfile, error) {
	var err error
	var filebytes []byte

	if filebytes, err = ioutil.ReadFile(testfile); err != nil {
		return nil, err
	}

	s := new(specfile)
	err = yaml.Unmarshal(filebytes, s)

	return s, err
}

func suite(s *specfile) error {
	if !s.isVersionValid() {
		return errors.New("Test file version is not valid")
	}

	for _, testing := range s.Testing {
		baseURL := buildURL(testing)

		for _, endpoint := range testing.Endpoints {
			if strings.HasPrefix(endpoint.Path, "/") {
				endpoint.Path = string(endpoint.Path[1:])
			}

			testResult := "OK"
			printTestResult(baseURL.String(), testResult)
		}
	}

	return nil
}

func buildURL(t testsection) url.URL {
	return url.URL{
		Scheme: t.Protocol,
		Host:   fmt.Sprintf(`%s:%v`, t.Host, t.Port),
	}
}

func printTestResult(url string, result string) {
	testTime := time.Now().Format(time.RFC822Z)
	output := fmt.Sprintf(testoutput, testTime, url, result)
	fmt.Println(output)
}
