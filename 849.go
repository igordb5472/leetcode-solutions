package main

import "slices"

func maxDistToClosest(seats []int) int {
	maxDist, prev1 := 0, slices.Index(seats, 1)
	maxDist = max(maxDist, prev1)
	for i := prev1 + 1; i < len(seats); i++ {
		if seats[i] == 1 {
			maxDist = max(maxDist, (i-prev1)/2)
			prev1 = i
		}
	}
	maxDist = max(maxDist, len(seats)-1-prev1)
	return maxDist
}
