package _0154__Find_Minimum_in_Rotated_Sorted_Array_II

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		switch {
		case nums[m] > nums[r]:
			l = m + 1
		case nums[m] < nums[r]:
			r = m
		default:
			r--
		}
	}
	return nums[l]
}
