package main

import "fmt"

func main() {
	arr := []int{8, 7, 9, 6, 1, 4}
	r := rob(arr)
	fmt.Print(r)
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	prev, cur := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		nums[i] += prev
		prev = max(prev, cur)
		cur = nums[i]
	}
	return max(prev, cur)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
