package main

func moveZeroes(nums []int) {
	for i, zeroCnt := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCnt++
			continue
		}
		if zeroCnt > 0 {
			nums[i-zeroCnt] ^= nums[i]
			nums[i] ^= nums[i-zeroCnt]
			nums[i-zeroCnt] ^= nums[i]
		}
	}
}
