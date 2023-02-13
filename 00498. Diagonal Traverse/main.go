package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := findDiagonalOrder(arr)
	fmt.Println(r)
}

func findDiagonalOrder(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	out := make([]int, n*m)
	col, row := 0, 0

	for i := 0; i < n*m; i++ {
		out[i] = mat[row][col]
		if (row+col)%2 == 0 {
			if col == m-1 {
				row++
			} else if row == 0 {
				col++
			} else {
				row, col = row-1, col+1
			}
		} else {
			if row == n-1 {
				col++
			} else if col == 0 {
				row++
			} else {
				row, col = row+1, col-1
			}
		}
	}

	return out
}

func findDiagonalOrder1(mat [][]int) []int {
	n, m := len(mat), len(mat[0])
	out := make([]int, n*m)

	t := true
	i, j := 0, 0
	for c := 0; c < n*m; c++ {
		out[c] = mat[i][j]
		if t {
			switch {
			case i == 0 && j == m-1:
				i, t = i+1, false
			case i != 0 && j != m-1:
				i, j = i-1, j+1
			case i == 0 && j != m-1:
				j, t = j+1, false
			case i != 0 && j == m-1:
				i, t = i+1, false
			}
		} else {
			switch {
			case j == 0 && i == n-1:
				j, t = j+1, true
			case j != 0 && i != n-1:
				j, i = j-1, i+1
			case j == 0 && i != n-1:
				i, t = i+1, true
			case j != 0 && i == n-1:
				j, t = j+1, true
			}
		}
	}

	return out
}
