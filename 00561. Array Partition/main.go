package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrayPairSum([]int{1, 4, 3, 2}))
}

func arrayPairSum(nums []int) int {
	if len(nums)%2 != 0 {
		return -1
	}

	sort.Ints(nums)

	out := 0
	for i := 0; i < len(nums); i = i + 2 {
		out += nums[i]
	}
	return out
}
