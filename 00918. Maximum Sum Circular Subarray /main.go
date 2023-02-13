package main

import "fmt"

func main() {
	arr := []int{1, -2, 3, -2}
	out := maxSubarraySumCircular(arr)
	fmt.Println(out)
}

func maxSubarraySumCircular(nums []int) int {
	sum, tmpMax, tmpMin := 0, 0, 0
	maxSum, minSum := nums[0], nums[0]
	for _, v := range nums {
		sum += v
		tmpMax = max(v, tmpMax+v)
		tmpMin = min(v, tmpMin+v)
		maxSum = max(tmpMax, maxSum)
		minSum = min(tmpMin, minSum)
	}
	if sum == minSum {
		return maxSum
	}
	return max(maxSum, sum-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxSubarraySumCircular1(nums []int) int {
	out := nums[0]
	n := len(nums)
	for i := 0; i < n; i++ {
		curMax := nums[i]
		for j := i + 1; j%n != i; j++ {
			curMax = max(curMax+nums[j%n], nums[j%n])
			out = max(out, curMax)
		}
	}
	return out
}
