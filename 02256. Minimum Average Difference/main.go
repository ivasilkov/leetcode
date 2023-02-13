package main

import (
	"fmt"
	"math"
)

func main() {
	//arr := []int{4, 2, 0}
	arr := []int{2, 5, 3, 9, 5, 3}
	r := minimumAverageDifference(arr)
	fmt.Println(r)
}

func minimumAverageDifference(nums []int) int {
	sum, n := 0, len(nums)
	for _, v := range nums {
		sum += v
	}

	minAvg, idx := math.MaxInt, 0
	leftSum, rightSum := 0, sum

	for i, v := range nums[:n-1] {
		leftSum += v
		rightSum -= v
		avgDiff := abs(leftSum/(i+1) - rightSum/(n-1-i))
		if avgDiff < minAvg {
			minAvg = avgDiff
			idx = i
		}
	}

	avgDiff := sum / n
	if avgDiff < minAvg {
		return n - 1
	}
	return idx
}

func minimumAverageDifference2(nums []int) int {
	n := len(nums)
	sums := make([]int, n)
	sums[0] = nums[0]
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	minAvg, idx := math.MaxInt, 0
	for i := 0; i < n-1; i++ {
		avgDiff := abs((sums[i] / (i + 1)) - ((sums[n-1] - sums[i]) / (n - i - 1)))
		if avgDiff < minAvg {
			minAvg = avgDiff
			idx = i
		}
	}

	avgDiff := sums[n-1] / n
	if avgDiff < minAvg {
		return n - 1
	}
	return idx
}

func minimumAverageDifference1(nums []int) int {
	min := math.MaxInt
	idx := 0

	for i := 1; i <= len(nums); i++ {
		firstSum := 0
		for _, v := range nums[:i] {
			firstSum += v
		}
		secondSum := 0
		for _, v := range nums[i:] {
			secondSum += v
		}
		var secondAvg int
		if i == len(nums) {
			secondAvg = 0
		} else {
			secondAvg = secondSum / (len(nums) - i)
		}
		avgDiff := abs((firstSum / i) - secondAvg)
		if avgDiff < min {
			min = avgDiff
			idx = i - 1
		}
	}

	return idx
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
