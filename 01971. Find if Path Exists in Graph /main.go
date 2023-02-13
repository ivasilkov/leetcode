package main

import "fmt"

func main() {
	r := validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2)
	fmt.Println(r)
}

// dfs
func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	m := make([][]int, n)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}

	st := []int{source}
	visited := make([]bool, n)
	visited[source] = true

	for len(st) != 0 {
		v := st[len(st)-1]
		st = st[:len(st)-1]

		for _, i := range m[v] {
			if i == destination {
				return true
			}
			if !visited[i] {
				visited[i] = true
				st = append(st, i)
			}
		}
	}
	return false
}

// bfs
func validPath1(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	m := make([][]int, n)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}

	q := []int{source}
	visited := make([]bool, n)
	visited[source] = true
	for len(q) != 0 {
		v := q[0]
		q = q[1:]
		for _, i := range m[v] {
			if i == destination {
				return true
			}
			if !visited[i] {
				q = append(q, i)
				visited[i] = true
			}
		}
	}
	return false
}
