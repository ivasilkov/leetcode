package main

import "fmt"

func main() {
	g := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	out := allPathsSourceTarget(g)
	fmt.Println(out)
}

func allPathsSourceTarget(graph [][]int) [][]int {
	last := len(graph) - 1

	var out [][]int
	var traverse func([]int, int)
	traverse = func(path []int, idx int) {
		path = append(path, idx)
		if idx == last {
			item := make([]int, len(path))
			copy(item, path)
			out = append(out, item)
		}
		for _, next := range graph[idx] {
			traverse(path, next)
		}
	}
	traverse([]int{}, 0)
	return out
}

func dfs(graph [][]int, idx int, m map[int][][]int) [][]int {
	if idx == len(graph)-1 {
		return [][]int{{idx}}
	}

	var out [][]int
	for _, nextIdx := range graph[idx] {
		paths, ok := m[nextIdx]
		if !ok {
			paths = dfs(graph, nextIdx, m)
			m[nextIdx] = paths
		}
		for _, path := range paths {
			out = append(out, append([]int{idx}, path...))
		}
	}
	return out
}
