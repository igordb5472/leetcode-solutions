package main

//TODO: fix [1, 0, ...]
func longestSubarray(nums []int) int {
	l, maxLen, cnt0 := 0, 0, (nums[0]+1)%2
	for r := 1; r < len(nums); r++ {
		if nums[r] == 0 {
			cnt0++
		}
		for cnt0 > 1 {
			if nums[l] == 0 {
				cnt0--
			}
			l++
		}
		maxLen = max(maxLen, r-l)
	}
	if cnt0 == 0 {
		return len(nums) - 1
	}
	cnt0++
	for cnt0 > 1 {
		if nums[l] == 0 {
			cnt0--
		}
		l++
	}
	return max(maxLen, len(nums)-l)
}
