package dailygogrill

func validParentheses(input string) bool {

	mapOfParentheses := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	parentheses := make([]string, len(input))

	for _, c := range input {
		_, exits := mapOfParentheses[string(c)]
		if exits {
			parentheses = append(parentheses, string(c))
		} else {
			lastParentheses := popLast(&parentheses)
			if mapOfParentheses[lastParentheses] != string(c) {
				return false
			}

		}
	}

	return true
}

func popLast(input *[]string) string {
	lastElement := (*input)[len(*input)-1]
	*input = (*input)[:len(*input)-1]
	return lastElement
}
