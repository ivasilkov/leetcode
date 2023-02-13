package main

import "fmt"

func main() {
	fmt.Println(strStr("sa1dbutsad", "sad"))

	fmt.Println(strStr("mississippi", "issipi"))
}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) || len(needle) == 0 || len(haystack) == 0 {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if needle == haystack[i:i+len(needle)] {
			return i
		}
	}

	return -1
}

func strStr1(haystack string, needle string) int {
	if len(needle) > len(haystack) || len(needle) == 0 || len(haystack) == 0 {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}

			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
