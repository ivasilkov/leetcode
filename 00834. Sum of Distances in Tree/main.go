package main

import "fmt"

func main() {
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	n := 6
	fmt.Println(sumOfDistancesInTree(n, edges))
}

// dfs
func sumOfDistancesInTree2(n int, edges [][]int) []int {
	res := make([]int, n)
	count := make([]int, n)
	tree := make([][]int, n)
	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}
	dfs(0, -1, tree, res, count)
	dfs2(0, -1, tree, res, count)
	return res
}

func dfs(root int, pre int, tree [][]int, res []int, count []int) {
	for _, c := range tree[root] {
		if c == pre {
			continue
		}
		dfs(c, root, tree, res, count)
		count[root] += count[c]
		res[root] += count[c] + res[c]
	}
	count[root]++
}

func dfs2(root int, pre int, tree [][]int, res []int, count []int) {
	for _, c := range tree[root] {
		if c == pre {
			continue
		}
		res[c] = res[root] - count[c]*2 + len(count)
		dfs2(c, root, tree, res, count)
	}
}

// bfs ???
type pos struct {
	x, y, next int
}

func sumOfDistancesInTree(n int, edges [][]int) []int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	q := make([]pos, 0, len(edges))
	graph := make([][]int, n)
	for _, edge := range edges {
		q = append(q, pos{edge[0], edge[1], 1})
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	for len(q) != 0 {
		v := q[0]
		q = q[1:]

		dp[v.x][v.y] = v.next
		dp[v.y][v.x] = v.next
		for _, node := range graph[v.y] {
			if dp[v.x][node] == 0 && v.x != node {
				q = append(q, pos{v.x, node, v.next + 1})
			}
		}
		for _, node := range graph[v.x] {
			if dp[v.y][node] == 0 && v.y != node {
				q = append(q, pos{v.y, node, v.next + 1})
			}
		}
	}

	out := make([]int, n)
	for i := range out {
		sum := 0
		for _, v := range dp[i] {
			sum += v
		}
		out[i] = sum
	}

	return out
}
