package main

func main() {

}

func longestPalindrome(s string) string {
	out := ""

	for i := range s {
		p1 := getPalindrome(s, i, i)
		p2 := getPalindrome(s, i, i+1)
		if len(out) < len(p1) {
			out = p1
		}
		if len(out) < len(p2) {
			out = p2
		}
	}

	return out
}

func getPalindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
