package main

import "fmt"

func main() {
	r := plusOne([]int{9, 9, 9})
	fmt.Println(r)
}

// in place
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	if digits[0] != 0 {
		return digits
	}

	n := make([]int, len(digits)+1)
	n[0] = 1
	return n
}
