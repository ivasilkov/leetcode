package main

import "fmt"

func main() {
	r := search([]int{-1, 0, 3, 5, 7, 9, 12}, 2)
	fmt.Println(r)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		case nums[mid] > target:
			right = mid - 1
		}
	}

	return -1
}
