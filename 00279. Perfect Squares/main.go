package main

import (
	"fmt"
)

func main() {
	r := numSquares(13)
	fmt.Println(r)
}

func numSquares(n int) int {
	q := []int{n}
	m := map[int]struct{}{}
	out := 0

	for len(q) > 0 {
		out++
		for _, v := range q {
			q = q[1:]
			for i := 1; i*i <= v; i++ {
				tmp := v - i*i
				if tmp == 0 {
					return out
				}
				if tmp < 0 {
					break
				}
				if _, ok := m[tmp]; ok {
					continue
				}
				q = append(q, tmp)
				m[tmp] = struct{}{}
			}
		}
	}

	return -1
}
