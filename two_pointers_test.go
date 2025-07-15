package dailygogrill

import "testing"

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{{
		input:    "Was it a car or a cat I saw?",
		expected: true,
	}}

	for _, test := range tests {
		actual := validPalindrome(test.input)
		if test.expected != actual {
			t.Errorf("Expected %v Actual %v", test.expected, actual)
		}
	}
}
