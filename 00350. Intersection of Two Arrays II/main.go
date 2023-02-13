package _0350__Intersection_of_Two_Arrays_II

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m := map[int]uint8{}
	for _, v := range nums1 {
		m[v]++
	}

	var out []int
	for _, v := range nums2 {
		if m[v] != 0 {
			out = append(out, v)
			m[v]--
		}
	}

	return out
}
