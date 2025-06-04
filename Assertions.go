package assertions

import (
	"reflect"
)

type TestReporter interface {
	Errorf(format string, args ...any)
}

func AssertInt(t TestReporter, expected, actual int, message string) {
	if expected != actual {
		t.Errorf("%s (expected: %v, actual: %v)", message, expected, actual)
	}
}

func AssertArray(t TestReporter, expected, actual []int, message string) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%s (expected: %v, actual: %v)", message, expected, actual)
	}
}
