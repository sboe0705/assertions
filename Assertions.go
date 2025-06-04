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

func AssertArray[T any](t TestReporter, expected, actual []T, message string) {
	if !reflect.DeepEqual(expected, actual) {
		if containsStructs(expected) {
			t.Errorf("%s (expected: %+v, actual: %+v)", message, expected, actual)
		} else {
			t.Errorf("%s (expected: %v, actual: %v)", message, expected, actual)
		}
	}
}

func containsStructs[T any](data []T) bool {
	for _, v := range data {
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() != reflect.Struct {
			return false
		}
	}
	return true
}
