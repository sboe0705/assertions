package assertions

import (
	"fmt"
	"strings"
	"testing"
)

type TestStruct struct {
	value int
}

type TestInterface any

func TestAssertTrue(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertTrue(mockT, true, "Values differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}

	// when
	AssertTrue(mockT, false, "Values differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Values differ (expected: true, actual: false)" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertFalse(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertFalse(mockT, false, "Values differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}

	// when
	AssertFalse(mockT, true, "Values differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Values differ (expected: false, actual: true)" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertEquals(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertEquals(mockT, 1, 1, "Values differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}

	// when
	AssertEquals(mockT, 1, 2, "Values differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if mockT.Messages[0] != "Values differ (expected: 1, actual: 2)" {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertEqualsWithInterface(t *testing.T) {
	mockT := &mockT{}

	object1 := &TestStruct{0}
	object2 := object1
	object3 := &TestStruct{0}

	// when
	AssertEquals(mockT, object1, object2, "Objects differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}

	// when
	AssertEquals(mockT, object1, object3, "Objects differ")

	// then
	if len(mockT.Messages) != 1 {
		t.Errorf("Expected 1 message")
	}

	if !strings.Contains(mockT.Messages[0], "Objects differ (expected: 0x") {
		t.Errorf("Expected different message than '%s'", mockT.Messages[0])
	}
}

func TestAssertArrayWithInt(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []int{1, 2, 3}, []int{1, 2, 3}, "Arrays differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}

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

func TestAssertArrayWithStruct(t *testing.T) {
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

func TestAssertArrayWithInterface(t *testing.T) {
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
