package main

import "slices"

func merge56(intervals [][]int) [][]int { // 56 resolves name conflicts
	merged := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int {
		switch {
		case a[0] < b[0] || a[0] == b[0] && a[1] < b[1]:
			return -1
		case a[0] > b[0] || a[0] == b[0] && a[1] > b[1]:
			return 1
		default:
			return 0
		}
	})
	l, r := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] <= r {
			r = max(r, interval[1])
			continue
		}
		merged = append(merged, []int{l, r})
		l, r = interval[0], interval[1]
	}
	if len(merged) == 0 || merged[len(merged)-1][1] != r {
		merged = append(merged, []int{l, r})
	}
	return merged
}
