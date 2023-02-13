package main

import "fmt"

func main() {
	r := pivotIndex([]int{-1, -1, -1, -1, -1, 0})
	fmt.Println(r)
}

func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftSum := 0
	for i, v := range nums {
		if 2*leftSum == sum-v {
			return i
		}
		leftSum += v
	}

	return -1
}
