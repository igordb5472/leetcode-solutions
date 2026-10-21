package main

func findClosestElements(arr []int, k int, x int) []int {
	var curDiffSum int
	for r := range k {
		if diff := x - arr[r]; diff >= 0 {
			curDiffSum += diff
		} else {
			curDiffSum -= diff
		}
	}
	start, end, minDiffSum := 0, k-1, curDiffSum
	for l, r := start+1, end+1; r < len(arr); l, r = l+1, r+1 {
		if diff := x - arr[l-1]; diff >= 0 {
			curDiffSum -= diff
		} else {
			curDiffSum += diff
		}
		if diff := x - arr[r]; diff >= 0 {
			curDiffSum += diff
		} else {
			curDiffSum -= diff
		}
		if curDiffSum < minDiffSum {
			start, end = l, r
			minDiffSum = curDiffSum
		}
	}
	return arr[start : end+1]
}
