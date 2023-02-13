package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}

// in-place and O(n) with space O(1)
func rotate(nums []int, k int) {
	if k < 0 {
		return
	}
	k = k % len(nums)

	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// in-place and O(n) with buf
func rotate1(nums []int, k int) {
	if k < 0 {
		return
	}

	l := len(nums)
	k = k % l
	buf := make([]int, k)
	for i := 0; i < k; i++ {
		buf[i] = nums[l-k+i]
	}

	for i := l - k - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}

	for i := 0; i < k; i++ {
		nums[i] = buf[i]
	}
}
