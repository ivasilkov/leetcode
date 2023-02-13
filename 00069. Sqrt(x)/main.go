package main

import (
	"fmt"
)

func main() {
	r := mySqrt(8)
	fmt.Println(r)
}

func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid

		if square == x {
			return mid
		}
		if square > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
