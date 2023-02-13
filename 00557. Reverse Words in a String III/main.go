package main

import "fmt"

func main() {
	fmt.Println(reverseWords3("Let's take LeetCode contest"))
}

func reverseWords3(s string) string {
	out := []byte(s)
	start := 0
	for i := 0; i < len(out); i++ {
		if out[i] != ' ' {
			continue
		}
		reverseBytes(out[start:i])
		start = i + 1
	}
	reverseBytes(out[start:len(out)])

	return string(out)
}

func reverseBytes(bytes []byte) {
	i, j := 0, len(bytes)-1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}
