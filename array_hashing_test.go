package dailygogrill

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	tests := []struct {
		input    []string
		expected [][]string
	}{{
		input:    []string{"act", "pots", "tops", "cat", "stop", "hat"},
		expected: [][]string{{"act", "cat"}, {"pots", "tops", "stop"}, {"hat"}},
	}}

	for _, ts := range tests {
		result := groupAnagrams(ts.input)
		if !reflect.DeepEqual(ts.expected, result) {
			t.Errorf("Expected %v Got %v", ts.expected, result)
		}
	}
}

func TestTopKFrequentElements(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{{
		input:    []int{1, 2, 2, 3, 3, 3},
		k:        2,
		expected: []int{3, 2},
	}}

	for _, test := range tests {
		got := topKFrequentElements(test.input, test.k)
		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("Expected %v, Got %v", test.expected, got)
		}
	}
}

func TestArrayHashing(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{{
		input:    []int{1, 2, 4, 6},
		expected: []int{48, 24, 12, 8},
	}}

	for _, test := range tests {
		got := producofArrayExceptSelf(test.input)

		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("Expected %v Got %v", test.expected, got)
		}
	}
}

func TestLongestConsecutiveSequence(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{{
		nums:     []int{2, 20, 4, 10, 3, 4, 5},
		expected: 4,
	}}

	for _, test := range tests {
		actual := longestConsecutiveSequence(test.nums)
		if !reflect.DeepEqual(test.expected, actual) {
			t.Errorf("Expected %d Actual %d", test.expected, actual)
		}
	}
}
