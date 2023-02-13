package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	fmt.Println(r)
}

// idea
// use int instead of [4]uint8
// use [10_000]bool instead of map

func openLock(deadends []string, target string) int {
	need, _ := strconv.Atoi(target)
	if need == 0 {
		return 0
	}

	visited := [10_000]bool{}
	for _, d := range deadends {
		i, _ := strconv.Atoi(d)
		visited[i] = true
	}
	if visited[0] {
		return -1
	}

	q := []int{0}
	visited[0] = true
	lvl := 0

	for len(q) != 0 {
		lvl++
		for _, v := range q {
			q = q[1:]
			for _, item := range getNextItems(v) {
				if item == need {
					return lvl
				}
				if visited[item] {
					continue
				}
				q = append(q, item)
				visited[item] = true
			}
		}
	}

	return -1
}

func getNextItems(num int) []int {
	out := make([]int, 0, 8)
	for i := 1; i <= 1_000; i *= 10 {
		v := (num / i) % 10
		if v == 0 {
			out = append(out, num+i, num+9*i)
		} else if v == 9 {
			out = append(out, num-9*i, num-i)
		} else {
			out = append(out, num+i, num-i)
		}
	}
	return out
}

/////////

func openLock1(deadends []string, target string) int {
	need := key(target)
	visited := map[[4]uint8]struct{}{}
	for _, v := range deadends {
		visited[key(v)] = struct{}{}
	}
	q := [][4]uint8{{0, 0, 0, 0}}
	if _, ok := visited[q[0]]; ok {
		return -1
	}
	visited[q[0]] = struct{}{}
	nextLevelIdx, lvl := 0, 0
	for idx := 0; idx < len(q); idx++ {
		v := q[idx]
		if v == need {
			return lvl
		}

		for i := 0; i < 4; i++ {
			newV := v
			if newV[i] == 9 {
				newV[i] = 0
			} else {
				newV[i]++
			}

			if _, ok := visited[newV]; !ok {
				q = append(q, newV)
				visited[newV] = struct{}{}
			}

			newV = v
			if newV[i] == 0 {
				newV[i] = 9
			} else {
				newV[i]--
			}
			if _, ok := visited[newV]; !ok {
				q = append(q, newV)
				visited[newV] = struct{}{}
			}
		}

		if idx == nextLevelIdx {
			lvl++
			nextLevelIdx = len(q) - 1
		}
	}

	return -1
}

func key(s string) [4]uint8 {
	out := [4]uint8{}
	for i := range s {
		out[i] = s[i] - '0'
	}
	return out
}
