package main

import "fmt"

func main() {
	fmt.Println(int('a' - 'a'))
	r := findMin([]int{3, 4, 5, 1, 2})
	fmt.Println(r)
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		m := (left + right) / 2
		if nums[m] > nums[right] {
			left = m + 1
		} else {
			right = m
		}
	}
	return nums[left]
}
