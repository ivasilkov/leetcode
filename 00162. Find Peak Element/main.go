package main

import "fmt"

func main() {
	r := findPeakElement([]int{1, 2, 3, 1})
	//r := findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})
	fmt.Println(r)
}

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] >= nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
