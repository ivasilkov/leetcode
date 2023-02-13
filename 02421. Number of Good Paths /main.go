package main

import (
	"fmt"
)

func main() {
	vals := []int{1, 3, 2, 1, 3}
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}
	r := numberOfGoodPaths(vals, edges)
	fmt.Println(r)
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	graph := make([][]int, len(vals))
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	out := 0
	for i := 0; i < len(vals); i++ {
		out += dfs(i, vals, i, -1, graph)
	}
	return out
}

func dfs(start int, vals []int, node, parent int, graph [][]int) int {
	if vals[node] > vals[start] {
		return 0
	}

	out := 0
	if vals[node] == vals[start] && start <= node {
		out++
	}

	for _, child := range graph[node] {
		if child == parent {
			continue
		}
		out += dfs(start, vals, child, node, graph)
	}
	return out
}

/// Union find
//
//type UnionFind struct {
//	root []int
//}
//
//func NewUF(n int) *UnionFind {
//	root := make([]int, n)
//	for i := 0; i < n; i++ {
//		root[i] = i
//	}
//	return &UnionFind{root}
//}
//
//func (u *UnionFind) Find(x int) int {
//	if u.root[x] != x {
//		u.root[x] = u.Find(u.root[x])
//	}
//	return u.root[x]
//}
//
//func (u *UnionFind) Union(x, y int) {
//	x, y = u.Find(x), u.Find(y)
//	u.root[x] = y
//}
//
//func numberOfGoodPaths(vals []int, edges [][]int) int {
//	n := len(vals)
//
//	val2Node := make(map[int][]int, n)
//	for i, val := range vals {
//		val2Node[val] = append(val2Node[val], i)
//	}
//
//	val2Edge := make(map[int][][]int, n)
//	for _, edge := range edges {
//		j, k := edge[0], edge[1]
//		maxVal := max(vals[j], vals[k])
//		val2Edge[maxVal] = append(val2Edge[maxVal], []int{j, k})
//	}
//
//	uf := NewUF(n)
//	res := 0
//
//	sortedUniqueVals := make([]int, 0, len(val2Node))
//	for k, _ := range val2Node {
//		sortedUniqueVals = append(sortedUniqueVals, k)
//	}
//
//	sort.Slice(sortedUniqueVals, func(i, j int) bool {
//		return sortedUniqueVals[i] < sortedUniqueVals[j]
//	})
//
//	for _, val := range sortedUniqueVals {
//		for _, edge := range val2Edge[val] {
//			uf.Union(edge[0], edge[1])
//		}
//		count := make(map[int]int)
//		for _, node := range val2Node[val] {
//			count[uf.Find(node)]++
//		}
//		for _, groupCount := range count {
//			res += int((groupCount * (groupCount + 1)) / 2)
//		}
//	}
//
//	return res
//}
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
