package main

import (
	"fmt"
)

func main() {
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	result := solutionESlice(s1, s2)
	fmt.Println(result)
}

func solutionESlice(s1, s2 string) int {
	if len(s1) != len(s2) {
		return 0
	}

	m := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		m[s1[i]-'a']++
		m[s2[i]-'a']--
	}

	for _, v := range m {
		if v != 0 {
			return 0
		}
	}
	return 1
}
