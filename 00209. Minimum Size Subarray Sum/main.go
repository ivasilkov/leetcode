package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))                               // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                                        // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))                        // 0
	fmt.Println(minSubArrayLen(213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12})) // 8
}

func minSubArrayLen(target int, nums []int) int {
	slow, fast, sum, min := 0, 0, 0, len(nums)+1
	for fast < len(nums) {
		for fast < len(nums) && sum < target {
			sum += nums[fast]
			fast++
		}
		for slow <= fast && sum >= target {
			min = minValue(min, fast-slow)
			sum -= nums[slow]
			slow++
		}
	}
	if min > len(nums) {
		return 0
	}
	return min
}

func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
