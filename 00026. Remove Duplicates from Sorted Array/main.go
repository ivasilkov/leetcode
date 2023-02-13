package _0026__Remove_Duplicates_from_Sorted_Array

func removeDuplicates(nums []int) int {
	writePtr := 1
	for readPtr := 1; readPtr < len(nums); readPtr++ {
		if nums[readPtr] == nums[readPtr-1] {
			continue
		}
		nums[writePtr] = nums[readPtr]
		writePtr++
	}
	return writePtr
}
