package main

import "fmt"

func main() {
	r := isPerfectSquare(16)
	fmt.Println(r)
}

func isPerfectSquare(num int) bool {
	l, r := 0, num
	for l <= r {
		m := (l + r) / 2
		square := m * m
		if square == num {
			return true
		}
		if square > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}
