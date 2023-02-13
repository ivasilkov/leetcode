package main

import "fmt"

func main() {
	r := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(r)
}

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minLen {
			minLen = len(strs[i])
		}
	}

	result := make([]byte, 0, minLen)

loop:
	for j := 0; j < minLen; j++ {
		for i := 1; i < len(strs); i++ {
			if strs[0][j] != strs[i][j] {
				break loop
			}
		}
		result = append(result, strs[0][j])
	}
	return string(result)
}
