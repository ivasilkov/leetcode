package main

import "fmt"

func main() {
	r := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(r)
}

func lengthOfLongestSubstring(s string) int {
	m, out := map[byte]int{}, 0
	var start int
	for i := 0; i < len(s); i++ {
		v, ok := m[s[i]]
		if ok {
			out = max(out, i-start)
			if v >= start {
				start = v + 1
			}
		}
		m[s[i]] = i
	}
	return max(out, len(s)-start)
}

func lengthOfLongestSubstring1(s string) int {
	m, out := map[byte]int{}, 0
	slow, fast := 0, 0
	for fast < len(s) {
		v, ok := m[s[fast]]
		if !ok {
			m[s[fast]] = fast
			fast++
			continue
		}
		out = max(out, fast-slow)
		fast = v + 1
		slow = fast
	}
	return out
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
