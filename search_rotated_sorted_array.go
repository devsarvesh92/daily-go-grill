package dailygogrill

func SearchRotatedSortedArray(numbers []int, target int) int {

	left, right := 0, len(numbers)-1

	for left < right {

		middle := (left + right) //2

		if numbers[middle] == target {
			return middle
		}

		if numbers[left] <= numbers[middle] {
			if target > numbers[left] && target < numbers[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target < numbers[right] && target > numbers[middle] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}
