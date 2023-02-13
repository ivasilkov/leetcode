package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("a"))
}

func removeDuplicates(s string) string {
	out, wPtr := []byte(s), 1
	for rPtr := 1; rPtr < len(s); rPtr++ {
		if wPtr != 0 && out[rPtr] == out[wPtr-1] {
			wPtr--
			continue
		}
		out[wPtr] = out[rPtr]
		wPtr++
	}
	return string(out[:wPtr])
}
