package assertions

import (
	"fmt"
	"testing"
)

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

func TestAssertArraySucceeding(t *testing.T) {
	mockT := &mockT{}

	// when
	AssertArray(mockT, []int{1, 2, 3}, []int{1, 2, 3}, "Arrays differ")

	// then
	if len(mockT.Messages) != 0 {
		t.Errorf("Expected no messages")
	}
}

func TestAssertArrayFailing(t *testing.T) {
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

// mocking

type mockT struct {
	Messages []string
}

func (m *mockT) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	m.Messages = append(m.Messages, msg)
}
