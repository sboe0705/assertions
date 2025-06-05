package assertions

import (
	"fmt"
	"strings"
	"testing"
)

type TestStruct struct {
	value int
}

type TestInterface interface {
}

func TestAssertIntSucceeding(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertInt(mockT, 1, 1, "Values differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}
}

func TestAssertIntFailing(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertInt(mockT, 1, 2, "Values differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Values differ (expected: 1, actual: 2)" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertObjectSucceeding(t *testing.T) {
	mockT := &mockT{}

	object1 := &TestStruct{0}
	object2 := object1

	// when
	AssertObject(mockT, object1, object2, "Objects differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}
}

func TestAssertObjectFailing(t *testing.T) {
	mockT := &mockT{}

	object1 := &TestStruct{0}
	object2 := &TestStruct{0}

	// when
	AssertObject(mockT, object1, object2, "Objects differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if !strings.Contains(mockT.Messages[0], "Objects differ") {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertIntArraySucceeding(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []int{1, 2, 3}, []int{1, 2, 3}, "Arrays differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}
}

func TestAssertIntArrayFailing(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []int{1, 2, 3}, []int{1, 2, 3, 4}, "Arrays differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Arrays differ (expected: [1 2 3], actual: [1 2 3 4])" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertStructArrayFailing(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []TestStruct{TestStruct{0}}, []TestStruct{TestStruct{1}}, "Arrays differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Arrays differ (expected: [{value:0}], actual: [{value:1}])" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertInterfaceArrayFailing(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []TestInterface{&TestStruct{0}}, []TestInterface{&TestStruct{1}}, "Arrays differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if !strings.Contains(mockT.Messages[0], "Arrays differ") {
		t.Errorf("Message '%s' missing required substring", mockT.Messages[0])
	}
}

// mocking

type mockT struct {
	Messages []string
}

func (m *mockT) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	m.Messages = append(m.Messages, msg)
}
