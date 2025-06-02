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
