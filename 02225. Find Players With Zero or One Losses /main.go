package _2225__Find_Players_With_Zero_or_One_Losses_

import "sort"

func findWinners(matches [][]int) [][]int {
	m := make(map[int]int)
	for _, i := range matches {
		m[i[0]] += 0
		m[i[1]]++
	}

	ans := [][]int{{}, {}}
	for k, v := range m {
		switch v {
		case 0:
			ans[0] = append(ans[0], k)
		case 1:
			ans[1] = append(ans[1], k)
		}
	}
	sort.Ints(ans[0])
	sort.Ints(ans[1])
	return ans
}
