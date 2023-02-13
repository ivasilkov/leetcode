package main

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(m)
}

func rotate(matrix [][]int) {
	l := len(matrix) - 1
	m := len(matrix) / 2

	for i := 0; i < m; i++ {
		for j := i; j < l-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[l-j][i]
			matrix[l-j][i] = matrix[l-i][l-j]
			matrix[l-i][l-j] = matrix[j][l-i]
			matrix[j][l-i] = tmp
		}
	}
}
