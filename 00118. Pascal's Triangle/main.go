package main

import (
	"fmt"
)

func main() {
	r := generate(7)
	fmt.Println(r)
}

func generate(numRows int) [][]int {
	if numRows < 1 {
		return [][]int{}
	}

	out := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		out[i] = make([]int, i+1)
		out[i][0], out[i][i] = 1, 1

		for j := 1; j < (i/2)+1; j++ {
			out[i][j] = out[i-1][j-1] + out[i-1][j]
			out[i][i-j] = out[i][j]
		}
	}

	return out
}
