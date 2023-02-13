package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	if n == 0 {
		return
	}

	var prev, cur string
	fmt.Scanf("%s", &prev)
	fmt.Println(prev)

	for i := 1; i < n; i++ {
		fmt.Scanf("%s", &cur)
		if cur == prev {
			continue
		}
		fmt.Println(cur)
		prev = cur
	}
}
