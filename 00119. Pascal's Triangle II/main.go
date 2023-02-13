package main

import "fmt"

func main() {
	fmt.Println(getRow(5))
}

// by formula f(i) = f(i-1) * (n - i + 1) / i
func getRow(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	out := make([]int, rowIndex+1)
	out[0], out[rowIndex] = 1, 1
	for i := 1; i <= rowIndex/2; i++ {
		out[i] = out[i-1] * (rowIndex - i + 1) / i
		out[rowIndex-i] = out[i]
	}
	return out
}

func getRow1(rowIndex int) []int {
	if rowIndex < 0 {
		return []int{}
	}

	out := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex/2; i++ {
		out[i] = calcPascalNum(rowIndex, i)
		out[rowIndex-i] = out[i]
	}
	return out
}

func calcPascalNum(rowIdx, idx int) int {
	if idx == 0 || idx == rowIdx {
		return 1
	}
	if idx == 1 {
		return rowIdx
	}
	return calcPascalNum(rowIdx-1, idx) + calcPascalNum(rowIdx-1, idx-1)
}
