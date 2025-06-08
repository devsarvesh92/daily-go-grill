package dailygogrill

import "sort"

func groupAnagrams(input []string) (result [][]string) {
	tracker := make(map[string][]string)

	for _, ip := range input {
		tracker[sortString(ip)] = append(tracker[sortString(ip)], ip)
	}

	for _, v := range tracker {
		result = append(result, v)
	}
	return
}

func sortString(ip string) string {
	runes := []rune(ip)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

type KeyValue struct {
	key   int
	value int
}

func topKFrequentElements(nums []int, k int) (result []int) {
	countTracker := make(map[int]int)

	for _, num := range nums {
		countTracker[num] += 1
	}
	pairs := sortDict(countTracker, true)

	ctr := 0
	for ctr < k {
		result = append(result, pairs[ctr].key)
		ctr += 1
	}
	return
}

func sortDict(dt map[int]int, descending bool) (pairs []KeyValue) {
	for k, v := range dt {
		pairs = append(pairs, KeyValue{key: k, value: v})
	}

	if descending {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].value > pairs[j].value
		})
	} else {
		sort.Slice(pairs, func(i, j int) bool {
			return pairs[i].value < pairs[j].value
		})
	}

	return
}
