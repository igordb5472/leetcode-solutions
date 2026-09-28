package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	ranges := make([]string, 0)
	left := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if left < nums[i-1] {
				ranges = append(ranges, fmt.Sprintf("%d->%d", left, nums[i-1]))
			} else {
				ranges = append(ranges, fmt.Sprintf("%d", left))
			}
			left = nums[i]
		}
	}
	if left < nums[len(nums)-1] {
		ranges = append(ranges, fmt.Sprintf("%d->%d", left, nums[len(nums)-1]))
	} else {
		ranges = append(ranges, fmt.Sprintf("%d", left))
	}
	return ranges
}
