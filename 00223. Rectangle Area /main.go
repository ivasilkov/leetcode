package main

import (
	"fmt"
)

func main() {
	r := computeArea(-2, -2, 2, 2, -2, -2, 2, 2)
	fmt.Println(r)
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)

	if ax1 >= bx2 || bx1 >= ax2 || ay1 >= by2 || by1 >= ay2 {
		// no intersection
		return area1 + area2
	}
	x1, y1 := max(ax1, bx1), max(ay1, by1)
	x2, y2 := min(ax2, bx2), min(ay2, by2)

	return area1 + area2 - ((x2 - x1) * (y2 - y1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
