package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	result := solutionB(arr)
	fmt.Println(result)
}

func solutionB(arr []int) int {
	out, cur := 0, 0
	for _, i := range arr {
		if i == 1 {
			cur++
		} else {
			out = max(cur, out)
			cur = 0
		}
	}
	return max(cur, out)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
