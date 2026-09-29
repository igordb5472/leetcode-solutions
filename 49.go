package main

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		groups[string(b)] = append(groups[string(b)], str)
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	return result
}
