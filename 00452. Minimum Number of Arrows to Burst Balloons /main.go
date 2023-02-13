package main

import "sort"

func main() {

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	out, end := len(points), points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			end = points[i][1]
		} else {
			out--
			if end > points[i][1] {
				end = points[i][1]
			}
		}
	}
	return out
}

func findMinArrowShots1(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	right, arrows := points[0][1], 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= right {
			right = min(right, points[i][1])
			continue
		} else {
			right = points[i][1]
			arrows++
		}
	}
	return arrows
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
