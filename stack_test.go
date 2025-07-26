package dailygogrill

import "testing"

func TestValidParentheses(t *testing.T) {
	testData := []struct {
		expected bool
		input    string
	}{{
		expected: true,
		input:    "([{}])",
	}}

	for _, testData := range testData {
		got := validParentheses(testData.input)
		if testData.expected != got {
			t.Errorf("Expected %v Got %v", testData.expected, got)
		}
	}
}
