package dailygogrill

import (
	"reflect"
	"testing"
)

func TestSearchRotatedSortedArray(t *testing.T) {

	testdata := []struct {
		numbers  []int
		target   int
		expected int
	}{{
		numbers:  []int{3, 4, 5, 6, 1, 2},
		target:   1,
		expected: 4,
	}}

	for _, test := range testdata {

		actual := SearchRotatedSortedArray(test.numbers, test.target)

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Expected %v actual %v", test.expected, actual)
		}
	}
}
