package main

func longestOnes(nums []int, k int) int {
	l, cnt0, maxLen := 0, (nums[0]+1)%2, 0
	for r := 1; r < len(nums); r++ {
		for cnt0 > k {
			cnt0 -= (nums[l] + 1) % 2
			l++
		}
		maxLen = max(maxLen, r-l)
		cnt0 += (nums[r] + 1) % 2
	}
	for cnt0 > k {
		cnt0 -= (nums[l] + 1) % 2
		l++
	}
	maxLen = max(maxLen, len(nums)-l)
	return maxLen
}
