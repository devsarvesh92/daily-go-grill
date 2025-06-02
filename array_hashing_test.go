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
