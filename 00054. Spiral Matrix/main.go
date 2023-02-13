package main

import "fmt"

func main() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := spiralOrder(arr)
	fmt.Println(r)
}

func spiralOrder(matrix [][]int) []int {
	rowLen, colLen := len(matrix), len(matrix[0])
	length := rowLen * colLen
	out := make([]int, length)
	i := 0

	rowStart, colStart, rowEnd, colEnd := 0, 0, rowLen-1, colLen-1

	for i < length {
		for idx := colStart; idx <= colEnd && i < length; idx++ {
			out[i] = matrix[rowStart][idx]
			i++
		}
		rowStart++

		for idx := rowStart; idx <= rowEnd && i < length; idx++ {
			out[i] = matrix[idx][colEnd]
			i++
		}
		colEnd--

		for idx := colEnd; idx >= colStart && i < length; idx-- {
			out[i] = matrix[rowEnd][idx]
			i++
		}
		rowEnd--

		for idx := rowEnd; idx >= rowStart && i < length; idx-- {
			out[i] = matrix[idx][colStart]
			i++
		}
		colStart++
	}

	return out
}
