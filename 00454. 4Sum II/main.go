package _0454__4Sum_II

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := map[int]int{}
	for _, i := range nums1 {
		for _, j := range nums2 {
			m[i+j]++
		}
	}
	out := 0
	for _, i := range nums3 {
		for _, j := range nums4 {
			out += m[-i-j]
		}
	}
	return out
}
