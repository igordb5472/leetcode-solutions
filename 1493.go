package main

//TODO: fix [1, 0, ...]
func longestSubarray(nums []int) int {
	maxLen, first0, second0 := 0, -1, -1
	for i := range nums {
		if nums[i] == 0 {
			if first0 != -1 && second0 != -1 {
				maxLen = max(maxLen, i-first0-2)
			}
			first0 = second0
			second0 = i
		}
	}
	if first0 == -1 {
		if second0 == -1 {
			return len(nums) - 1
		}
		if second0-first0 == 1 {
			return max(first0, len(nums)-1-second0)
		}
		return max(second0-1, len(nums)-2-first0)
	}
	return max(maxLen, len(nums)-first0-2)
}
