package dailygogrill_test

import (
	"dailygogrill"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	testData := []struct {
		input    []int
		number   int
		expected int
	}{{
		input:    []int{3, 4, 5, 6},
		number:   4,
		expected: 1,
	}}

	for _, val := range testData {
		got := dailygogrill.FindNumber(val.input, val.number)
		if got != val.expected {
			t.Errorf("Expected %v got %v", got, val.expected)
		}
	}
}
