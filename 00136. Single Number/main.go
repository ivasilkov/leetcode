package main

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// hashset
func singleNumber1(nums []int) int {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = struct{}{}
		}
	}
	for k := range m {
		return k
	}
	return -1
}
