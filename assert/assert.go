package assert

import (
	"net/http"

	"github.com/eclesiomelojunior/gospec/apispec"
)

// Assert struct
type assert struct {
}

// NewAssert func
func NewAssert() apispec.Assert {
	return &assert{}
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
