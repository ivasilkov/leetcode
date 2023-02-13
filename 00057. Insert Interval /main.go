package main

import "fmt"

func main() {
	newInterval := []int{4, 5}
	intervals := [][]int{{2, 3}, {6, 9}}
	out := insert(intervals, newInterval)
	fmt.Println(out)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	start := newInterval[0]
	end := newInterval[1]
	for _, interval := range intervals {
		switch {
		case interval[1] < newInterval[0]:
			res = append(res, interval)
		case interval[0] > newInterval[1]:
			if start != -1 {
				res = append(res, []int{start, end})
				start = -1
			}
			res = append(res, interval)
		case interval[0] < newInterval[0]:
			start = interval[0]
			if interval[1] > newInterval[1] {
				end = interval[1]
			}
		case interval[1] > newInterval[1]:
			end = interval[1]
		}
	}
	if start != -1 {
		res = append(res, []int{start, end})
	}
	return res
}
