package main

import "fmt"

func main() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	labels := "abaedcd"
	ans := countSubTrees(n, edges, labels)
	fmt.Println(ans)
}

func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	ans := make([]int, n)
	dfs(0, -1, labels, graph, ans)
	return ans
}

func dfs(node, parent int, labels string, graph [][]int, ans []int) [26]int {
	m := [26]int{}
	m[labels[node]-'a'] = 1

	for _, child := range graph[node] {
		if child == parent {
			continue
		}
		sub := dfs(child, node, labels, graph, ans)
		for i := range sub {
			m[i] += sub[i]
		}
	}

	ans[node] = m[labels[node]-'a']
	return m
}
