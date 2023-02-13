package main

import "fmt"

func main() {
	n := 4
	arr := [][]int{{1, 2}, {1, 3}, {2, 4}}
	r := possibleBipartition(n, arr)
	fmt.Println(r)
}

func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n+1)
	for _, dislike := range dislikes {
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}

	colors := make([]int8, n+1)
	for i := 1; i <= n; i++ {
		if colors[i] != 0 {
			continue
		}
		if !dfs(colors, graph, i, 1) {
			return false
		}
	}
	return true
}

func dfs(colors []int8, graph [][]int, node int, color int8) bool {
	colors[node] = color

	for _, connectedNode := range graph[node] {
		if colors[connectedNode] == color {
			return false
		}
		if colors[connectedNode] != 0 {
			continue
		}
		if !dfs(colors, graph, connectedNode, color*-1) {
			return false
		}
	}

	return true
}
