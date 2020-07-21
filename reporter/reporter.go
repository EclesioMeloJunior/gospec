package reporter

import "github.com/eclesiomelojunior/gospec/apispec"

//NewCLIReporter interface
func NewCLIReporter() apispec.Reporter {
	return &cliReporter{}
}

type cliReporter struct {
}

func (cli *cliReporter) Results([]byte) {
}
