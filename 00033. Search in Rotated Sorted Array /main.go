package main

import "fmt"

func main() {
	r := search([]int{3, 1}, 1)
	fmt.Println(r)
}

func search1(nums []int, target int) int {
	start := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			start = i
			break
		}
	}

	if start != 0 {
		nums = append(nums[start:], nums[:start]...)
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return (start + mid) % len(nums)
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/////////

func search(nums []int, target int) int {
	n := len(nums)

	l, r := 0, len(nums)
	for l <= r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	pivot := l
	l, r = pivot, pivot-1+n

	for l <= r {
		m := (l + r) / 2
		val := nums[m%n]
		if val == target {
			return m % n
		}
		if val > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
