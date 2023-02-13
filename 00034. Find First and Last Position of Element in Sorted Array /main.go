package main

import "fmt"

func main() {
	arr := []int{5, 7, 8, 8, 9, 10}
	target := 8
	r := searchRange(arr, target)
	fmt.Println(r)
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{leftBinary(nums, target), rightBinary(nums, target)}
}

func leftBinary(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func rightBinary(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l+r)/2 + 1
		if nums[m] > target {
			r = m - 1
		} else {
			l = m
		}
	}
	if nums[r] == target {
		return r
	}
	return -1
}

func searchRange1(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			break
		}
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	lr := (l + r) / 2
	rl := lr
	for l <= lr {
		m := (l + lr) / 2
		if nums[m] == target {
			break
		}
		if nums[m] > target {
			lr = m - 1
		} else {
			l = m + 1
		}
	}

	for rl <= r {
		m := (rl + r) / 2
		if nums[m] == target {
			break
		}
		if nums[m] > target {
			r = m - 1
		} else {
			rl = m + 1
		}
	}

	if nums[l] == target && nums[rl] == target {
		return []int{l, rl}
	}
	if nums[l] == target {
		return []int{l, l}
	}
	if nums[rl] == target {
		return []int{rl, rl}
	}
	return []int{-1, -1}
}
