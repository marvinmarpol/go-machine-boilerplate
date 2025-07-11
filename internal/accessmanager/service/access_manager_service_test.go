package service

import "testing"

func TestAddRole(t *testing.T) {
	service := NewAccessManagerService()

	testCases := []struct {
		input    string
		expected bool
	}{
		{"admin", true},
		{"viewer", true},
		{"admin", false},
	}

	for i, test := range testCases {
		result := service.AddRole(test.input)
		if result != test.expected {
			t.Errorf("\ntest number %d when creating %s role failed. Expected %v Got %v\n", (i + 1), test.input, test.expected, result)
		}
	}

}
