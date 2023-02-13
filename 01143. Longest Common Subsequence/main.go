package main

import (
	"fmt"
)

func main() {
	t1 := "bcdeca"
	t2 := "acec"
	r := longestCommonSubsequence(t1, t2)
	fmt.Println(r)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	l1, l2 := len(text1), len(text2)
	dp := make([][]int16, l1+1)
	for i := range dp {
		dp[i] = make([]int16, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maxInt16(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return int(dp[l1][l2])
}

func maxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	out := 0
	for i := 0; i < len(text1); i++ {
		cur := 0
		for j, k := i, 0; k < len(text2) && j < len(text1); k++ {
			if text1[j] == text2[k] {
				cur++
				j++
			}
		}
		out = max(out, cur)
	}
	return out
}

func longestCommonSubsequence1(text1 string, text2 string) int {
	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}

	out, cur, prev := 0, 0, -1
	for i := range text1 {
		v := text1[i]
		idx := find(text2, v, 0)
		if idx == -1 {
			out = max(out, cur)
			cur = 0
			prev = -1
			continue
		}
		idx = find(text2, v, prev)
		if idx < prev {
			out = max(out, cur)
			cur = 1
			prev = idx
			continue
		}
		prev = idx
		cur++
	}
	return max(out, cur)
}

func find(str string, b byte, start int) int {
	if start == -1 {
		start = 0
	}
	for i := start; i < len(str); i++ {
		if str[i] == b {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
