package _0587__Erect_the_Fence

import "sort"

func outerTrees(trees [][]int) [][]int {
	if len(trees) <= 3 {
		return trees
	}

	sort.Slice(trees, func(q, p int) bool {
		if trees[q][0]-trees[p][0] == 0 {
			return trees[q][1] < trees[p][1]
		}
		return trees[q][0] < trees[p][0]
	})

	upper, lower := make([][]int, 0), make([][]int, 0)
	for i := range trees {
		for len(lower) >= 2 && compare(lower[len(lower)-2], lower[len(lower)-1], trees[i]) > 0 {
			lower = lower[:len(lower)-1]
		}
		for len(upper) >= 2 && compare(upper[len(upper)-2], upper[len(upper)-1], trees[i]) < 0 {
			upper = upper[:len(upper)-1]
		}
		lower = append(lower, trees[i])
		upper = append(upper, trees[i])
	}

	result := make([][]int, 0)
	m := make(map[int]struct{})

	for _, set := range append(upper, lower...) {
		k := set[0]*1000 + set[1] // 0 <= x,y <= 100
		if _, ok := m[k]; ok {
			continue
		}
		result = append(result, set)
		m[k] = struct{}{}
	}

	return result
}

func compare(p, q, r []int) int {
	return (r[1]-q[1])*(q[0]-p[0]) - (q[1]-p[1])*(r[0]-q[0])
}
