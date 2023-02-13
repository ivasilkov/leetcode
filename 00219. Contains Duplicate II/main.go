package _0219__Contains_Duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := range nums {
		v, ok := m[nums[i]]
		if ok && i-v <= k { // i > v
			return true
		}
		m[nums[i]] = i
	}
	return false
}
