package _0387__First_Unique_Character_in_a_String

func firstUniqChar(s string) int {
	m := map[byte]byte{}
	for i := range s {
		m[s[i]] += 1
	}
	for i := range s {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar1(s string) int {
loop:
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				continue loop
			}
		}
		return i
	}
	return -1
}
