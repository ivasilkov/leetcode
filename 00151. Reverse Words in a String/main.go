package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1 2 3 4 -> 4 3 2 1
	fmt.Println(reverseWords("  the  sky  is  blue"))
}

func reverseWords(s string) string {
	// remove extra spaces
	words := strings.Split(strings.Trim(s, " "), " ")
	wPos := 0
	for rPos := 0; rPos < len(words); rPos++ {
		if words[rPos] == "" {
			continue
		}
		words[wPos] = words[rPos]
		wPos++
	}

	// reverse words
	i, j := 0, wPos-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}

	return strings.Join(words[:wPos], " ")
}

// reverse strings, reverse each word
func reverseWords1(s string) string {
	start, end := 0, len(s)-1
	for start < end {
		startIsSpace := s[start] == ' '
		endIsSpace := s[end] == ' '
		if !endIsSpace && !startIsSpace {
			break
		}
		if startIsSpace {
			start++
		}
		if endIsSpace {
			end--
		}
	}

	out := make([]byte, 0, end-start+1)
	for i := start; i <= end; i++ {
		if s[i] == ' ' && s[i-1] == ' ' {
			continue
		}
		out = append(out, s[i])
	}

	reverseBytes(out)
	start = 0
	for i := 0; i < len(out); i++ {
		if out[i] != ' ' {
			continue
		}
		reverseBytes(out[start:i])
		start = i + 1
	}
	reverseBytes(out[start:])

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
