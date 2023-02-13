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

	recursiveD("(", 1, 0, n)
}

func recursiveD(s string, op, cl, n int) {
	if op == n && cl == n {
		fmt.Println(s)
		return
	}

	if op < n {
		recursiveD(s+"(", op+1, cl, n)
	}

	if cl < n && op > cl {
		recursiveD(s+")", op, cl+1, n)
	}
}
