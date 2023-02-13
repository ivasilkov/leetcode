package main

import (
	"fmt"
	"math"
)

func main() {
	r := dominantIndex([]int{3, 6, 1, 0})
	fmt.Println(r)
}

func dominantIndex(nums []int) int {
	max, secondMax := math.MinInt64, math.MinInt64
	idx := -1
	for i, v := range nums {
		if v > max {
			secondMax, max = max, v
			idx = i
		} else if v > secondMax {
			secondMax = v
		}
	}

	if secondMax*2 > max {
		return -1
	}
	return idx
}

func dominantIndex1(nums []int) int {
	maxIdx := 0
	for i, v := range nums {
		if v > nums[maxIdx] {
			maxIdx = i
		}
	}

	for i, v := range nums {
		if v*2 > nums[maxIdx] && i != maxIdx {
			return -1
		}
	}
	return maxIdx
}
