package main

func rotate(matrix [][]int) {
	for first, last := 0, len(matrix)-1; first < last; first, last = first+1, last-1 {
		for i := first + 1; i <= last; i++ {
			matrix[first][i], matrix[i][last] = matrix[i][last], matrix[first][i]
		}
		for i := last - 1; i >= first; i-- {
			matrix[last][i], matrix[i][first] = matrix[i][first], matrix[last][i]
		}
		for i := first + 1; i <= last; i++ {
			matrix[first][i], matrix[last][len(matrix)-1-i] = matrix[last][len(matrix)-1-i], matrix[first][i]
		}
	}
}
