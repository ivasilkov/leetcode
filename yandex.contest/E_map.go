package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	result := solutionE(s1, s2)
	fmt.Println(result)
}

func solutionE(s1, s2 string) int {
	if len(s1) != len(s2) {
		return 0
	}

	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		m[r]--
	}
	for _, v := range m {
		if v != 0 {
			return 0
		}
	}
	return 1
}
