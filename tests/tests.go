package tests

import (
	"fmt"
	"testing"
)

// ExpectationError abstract the message of expectation erros at testing
func ExpectationError(t *testing.T, expected, received interface{}, message string) {
	expectationErrorMessage := fmt.Sprintf(
		"Expected: %s\nReceived: %s at %s",
		expected, received, message,
	)

	t.Error(expectationErrorMessage)
}
