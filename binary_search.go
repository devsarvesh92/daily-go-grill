package dailygogrill

import (
	"sort"
)

func FindNumber(numbers []int, num int) int {
	i := 0
	j := len(numbers) - 1

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	middle := 0

	for i < j {

		middle = (i + j) //2
		if numbers[middle] == num {
			return middle
		} else if num > numbers[middle] {
			i = middle + 1
		} else {
			j = middle - 1
		}
	}
	return middle
}
