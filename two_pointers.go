package dailygogrill

import (
	"strings"
	"unicode"
)

func validPalindrome(s string) bool {
	cleanString := ""

	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			cleanString += string(strings.ToLower(string(c)))
		}
	}

	start, end := 0, len(cleanString)-1

	for start < end {
		if cleanString[start] != cleanString[end] {
			return false
		}
		start += 1
		end -= 1
	}
	return true
}
