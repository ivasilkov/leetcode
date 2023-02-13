package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("badc", "baba"))
}

// hashset + hashmap
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mem, set := map[byte]byte{}, map[byte]struct{}{}
	for idx := range s {
		v, inMap := mem[s[idx]]
		_, inSet := set[t[idx]]
		if !inMap && !inSet {
			mem[s[idx]], set[t[idx]] = t[idx], struct{}{}
			continue
		}
		if v != t[idx] {
			return false
		}
	}
	return true
}

// double hashmap
func isIsomorphic1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := map[byte]int{}, map[byte]int{}
	for idx := range s {
		if sMap[s[idx]] != tMap[t[idx]] {
			return false
		}
		sMap[s[idx]] = idx + 1
		tMap[t[idx]] = sMap[s[idx]]
	}
	return true
}

// hashmap
func isIsomorphic2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[int]int{}
	for idx := range s {
		if m[int(s[idx])] != m[-1*int(t[idx])] {
			return false
		}
		m[int(s[idx])] = idx + 1
		m[-1*int(t[idx])] = idx + 1
	}
	return true
}
