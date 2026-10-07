package main

func missingNumber(nums []int) int {
	n := len(nums) * (len(nums) + 1) / 2
	for _, x := range nums {
		n -= x
	}
	return n
}
