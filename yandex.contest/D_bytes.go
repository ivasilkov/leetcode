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

	bytes := make([]byte, n*2)
	bytes[0] = '('
	recursiveDUseBytes(bytes, 1, 0, n)
}

func recursiveDUseBytes(bytes []byte, op, cl, n int) {
	if op == n && cl == n {
		fmt.Println(string(bytes))
		return
	}

	if op < n {
		bytes[op+cl] = '('
		recursiveDUseBytes(bytes, op+1, cl, n)
	}

	if cl < n && op > cl {
		bytes[op+cl] = ')'
		recursiveDUseBytes(bytes, op, cl+1, n)
	}
}
