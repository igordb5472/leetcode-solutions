package main

func trap(height []int) int {
	sum, l, r, leftMax, rightMax := 0, 0, len(height)-1, height[0], height[len(height)-1]
	for l < r {
		if height[l] < height[r] {
			if height[l] > leftMax {
				leftMax = height[l]
			} else {
				sum += leftMax - height[l]
			}
			l++
		} else {
			if height[r] > rightMax {
				rightMax = height[r]
			} else {
				sum += rightMax - height[r]
			}
			r++
		}
	}
	return sum
}
