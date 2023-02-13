package main

import (
	"fmt"
)

type coord struct {
	x, y int
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	coords := make([]coord, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &coords[i].x)
		fmt.Scanf("%d", &coords[i].y)
	}

	var k int
	fmt.Scanf("%d", &k)

	var start, end int
	fmt.Scanf("%d", &start)
	fmt.Scanf("%d", &end)

	if start == end {
		fmt.Println(0)
		return
	}
	start--
	end--

	graph := makeGraph(coords, n, k)
	result := solutionG(start, end, n, graph)
	fmt.Println(result)
}

func solutionG(start, end, n int, graph [][]int) int {
	q := []int{start}
	visited := make([]bool, n)
	visited[start] = true

	out := 0
	for len(q) > 0 {
		out++
		for _, v := range q {
			q = q[1:]
			for _, i := range graph[v] {
				if visited[i] {
					continue
				}
				if i == end {
					return out
				}
				q = append(q, i)
				visited[i] = true
			}
		}
	}

	return -1
}

func makeGraph(coords []coord, n, k int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := distance(coords[i], coords[j])
			if d > k {
				continue
			}
			graph[i] = append(graph[i], j)
			graph[j] = append(graph[j], i)
		}
	}
	return graph
}

func distance(c1, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
