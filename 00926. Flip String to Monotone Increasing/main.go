package main

func main() {

}

func minFlipsMonoIncr(s string) int {
	f0, f1 := 0, 0
	for i := 0; i < len(s); i++ {
		tmp := int(s[i] - '0')
		f0 += tmp
		f1 = min(f0, f1+1-tmp)
	}
	return f1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
