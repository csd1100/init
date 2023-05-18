package utils

const FAILURE_MESSAGE = "%s: assertion failed \nExpected:\n\t %v \nActual:\n\t %v"

type TestCase struct {
	Name           string
	args           []any
	expected_value any
	expected_error error
}
