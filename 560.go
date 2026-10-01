package main

func subarraySum(nums []int, k int) int {
	cnt, sum, freq := 0, 0, map[int]int{0: 1}
	for _, x := range nums {
		sum += x
		cnt += freq[sum-k]
		freq[sum]++
	}
	return cnt
}
