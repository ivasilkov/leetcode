package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := "aabbbabbbbbabbbbaabaabbaaabbbabaababaaaabbbbbabbaa"
	s2 := "aabbaabbbabaababaabaababbbababbbaaaabbbbbabbbaabaa"
	baseStr := "buqpqxmnajphtisernebttymtrydomxnwonfhfjlzzrfhosjct"
	r := smallestEquivalentString(s1, s2, baseStr)
	fmt.Println(r)
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	m := map[byte]int{}
	groups := [][]byte{}

	for i := 0; i < len(s1); i++ {
		char1, char2 := s1[i], s2[i]
		v1, ok1 := m[char1]
		v2, ok2 := m[char2]
		switch {
		case !ok1 && !ok2:
			// save items to new group
			m[char1] = len(groups)
			m[char2] = len(groups)
			groups = append(groups, []byte{char1, char2})
			continue
		case ok1 && ok2:
			if v1 == v2 {
				// already in same group
				continue
			}
			// merge groups v2 and v1
			groups[v1] = append(groups[v1], groups[v2]...)
			for _, char := range groups[v2] {
				m[char] = v1
			}
		case ok1: // and !ok2
			// and char2 to group v1
			m[char2] = v1
			groups[v1] = append(groups[v1], char2)
		case ok2: // and !ok1
			// and char1 to group v2
			m[char1] = v2
			groups[v2] = append(groups[v2], char1)
		}
	}

	// sort chars in each group
	for k := range groups {
		sort.Slice(groups[k], func(i, j int) bool {
			return groups[k][i] < groups[k][j]
		})
	}

	// make out string
	out := make([]byte, len(baseStr))
	for i := range baseStr {
		g, ok := m[baseStr[i]]
		if ok {
			out[i] = groups[g][0]
		} else {
			out[i] = baseStr[i]
		}
	}
	return string(out)
}
