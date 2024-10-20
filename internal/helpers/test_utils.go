package helpers

import (
	"errors"
	"testing"
)

const failureMessage = `%s: assertion failed when validating %s
	Expected:
		%v
	Actual:
		%v
	With Error: %v`

const errorAssertion = "errors are not same"

func ValidateExpectations(t *testing.T, name string, actualValue, expectedValue any, actualError, expectedError error, validationFn func(any, any) error) {

	if actualError != nil {
		if !errors.Is(actualError, expectedError) {
			t.Errorf(failureMessage, name, "error", expectedError, actualError, errorAssertion)
		}
	} else if expectedError != nil {
		t.Errorf(failureMessage, name, "error", expectedError, nil, errorAssertion)
	} else {
		valueErr := validationFn(expectedValue, actualValue)
		if valueErr != nil {
			t.Errorf(failureMessage, name, "value", expectedValue, actualValue, valueErr)
		}
	}
}
