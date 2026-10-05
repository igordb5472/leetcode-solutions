package main

func generateMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	left, right, top, bottom, cnt := 0, n-1, 0, n-1, 1
	for {
		for j := left; j <= right; j++ {
			m[top][j] = cnt
			cnt++
		}
		if cnt > n*n {
			break
		}
		top++
		for i := top; i <= bottom; i++ {
			m[i][right] = cnt
			cnt++
		}
		if cnt > n*n {
			break
		}
		right--
		for j := right; j >= left; j-- {
			m[bottom][j] = cnt
			cnt++
		}
		if cnt > n*n {
			break
		}
		bottom--
		for i := bottom; i >= top; i-- {
			m[i][left] = cnt
			cnt++
		}
		if cnt > n*n {
			break
		}
		left++
	}
	return m
}
