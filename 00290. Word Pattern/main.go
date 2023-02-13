package main

import "strings"

func main() {

}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	m := make(map[byte]string)
	uniq := make(map[string]struct{})

	for i := range pattern {
		v, ok := m[pattern[i]]
		if ok && words[i] != v {
			return false
		} else if !ok {
			m[pattern[i]] = words[i]
			if _, ok := uniq[words[i]]; ok {
				return false
			}
			uniq[words[i]] = struct{}{}
		}
	}
	return true
}
