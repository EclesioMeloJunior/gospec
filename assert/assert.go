package assert

import (
	"net/http"

	"github.com/eclesiomelojunior/gospec/apispec"
)

// Assert struct
type assert struct {
	reporter apispec.Reporter
}

// NewAssert func
func NewAssert(reporter apispec.Reporter) apispec.Assert {
	return &assert{
		reporter: reporter,
	}
}

// Body method
func (a *assert) TestBody(expected interface{}) apispec.Assert {
	return a
}

// Headers method
func (a *assert) TestHeaders(expected http.Header) apispec.Assert {
	return a
}

// StatusCode method
func (a *assert) TestStatusCode(expected int) apispec.Assert {
	return a
}

func (a *assert) Test(*http.Response) ([]byte, error) {
	return []byte{}, nil
}
