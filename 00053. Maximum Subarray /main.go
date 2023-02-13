package main

func main() {

}

func maxSubArray(nums []int) int {
	maxSum, tmpMax := nums[0], 0
	for _, v := range nums {
		tmpMax = max(v, tmpMax+v)
		maxSum = max(maxSum, tmpMax)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
