package main

func main() {

}

func canJump(nums []int) bool {
	pos, last := nums[0], len(nums)-1
	for i := 0; i <= pos; i++ {
		next := i + nums[i]
		if next > pos {
			pos = next
		}
		if pos >= last {
			return true
		}
	}
	return false
}
