package reporter

import "github.com/eclesiomelojunior/gospec/apispec"

//NewReporter interface
func NewReporter() apispec.Reporter {
	return &cliReporter{}
}

type cliReporter struct {
}

func (cli *cliReporter) Results([]byte) {
}
