package _1704__Determine_if_String_Halves_Are_Alike_

func halvesAreAlike(s string) bool {
	counter := 0
	for i := 0; i < len(s)/2; i++ {
		if isVowel(s[i]) {
			counter++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		if isVowel(s[i]) {
			counter--
		}
	}
	return counter == 0
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
