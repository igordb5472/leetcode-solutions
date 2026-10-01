package main

func twoSum(nums []int, target int) []int {
	prev := make(map[int]int)
	for idxX, x := range nums {
		if idxY, ok := prev[target-x]; ok {
			return []int{idxY, idxX}
		}
		prev[x] = idxX
	}
	return nil
}
