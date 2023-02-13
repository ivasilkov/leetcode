package main

func main() {

}

func detectCapitalUse(word string) bool {
	count := 0
	for i := range word {
		if isUpper(word[i]) {
			count++
		}
	}
	return count == len(word) ||
		count == 0 ||
		(count == 1 && isUpper(word[0]))
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func detectCapitalUse1(word string) bool {
	prev := true

	for i := range word {
		if word[i] >= 'A' && word[i] <= 'Z' {
			if !prev {
				return false
			}
			prev = true
		} else {
			if prev && i > 1 {
				return false
			}
			prev = false
		}
	}

	return true
}
