package mocking

import "testing"

// NOTE: Explanation:
/*
	•	Interface: Service defines a contract for what a service should do.
	•	Mock implementation: MockService is used for testing without depending on the real
		implementation.
	•	Testing: The test checks that UseService behaves correctly with the mock.
*/

type MockService struct{}

func (service MockService) DoSomething() string {
	return "Mock doing something"
}

func TestUseService(test *testing.T) {
	mock := MockService{}
	result := UseService(mock)
	expected := "Mock doing something"
	if result != expected {
		test.Errorf("expected %s, got %s", expected, result)
	}
}
