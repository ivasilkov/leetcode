package main

func intersection(nums1 []int, nums2 []int) []int {
	out, m := make([]int, 0), make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m[v]; ok {
			out = append(out, v)
			delete(m, v)
		}
	}
	return out
}
