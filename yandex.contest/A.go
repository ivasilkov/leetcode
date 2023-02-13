package main

import (
	"fmt"
)

func main() {
	var j, s string
	fmt.Scanf("%s", &j)
	fmt.Scanf("%s", &s)

	result := solutionA(j, s)
	fmt.Println(result)
}

func solutionA(j, s string) int {
	m := make(map[byte]struct{}, len(j))
	for i := range j {
		m[j[i]] = struct{}{}
	}

	out := 0
	for i := range s {
		_, ok := m[s[i]]
		if ok {
			out++
		}
	}
	return out
}
