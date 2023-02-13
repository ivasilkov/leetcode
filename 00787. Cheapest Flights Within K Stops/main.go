package main

import "math"

func main() {
	n := 4
	src := 0
	dst := 3
	k := 5

	r := findCheapestPrice()
}

type item struct {
	dst   int
	price int
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]item, n)
	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], item{flight[1], flight[2]})
	}

	out := make([]int, n)
	for i := range out {
		out[i] = math.MaxInt
	}
	out[src] = 0

	q := []int{src}
	for i := 0; len(q) > 0 && i < k; i++ {
		for _, v := range q {
			q = q[1:]
			for _, it := range graph[v] {
				out[it.dst] = min(out[v]+it.price, out[it.dst])
				q = append(q, it.dst)
			}
		}
	}

	return out[dst]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
	graph := make([][]item, n)
	for _, flight := range flights {
		graph[flight[0]] = append(graph[flight[0]], item{flight[1], flight[2]})
	}

	out := math.MaxInt
	dfs(graph, src, dst, k, 0, &out)
	if out == math.MaxInt {
		return -1
	}
	return out
}

func dfs(graph [][]item, src, dst, k, prev int, out *int) {
	if src == dst {
		*out = min(*out, prev)
		return
	}
	if k < 0 || prev > *out {
		return
	}
	for _, it := range graph[src] {
		dfs(graph, it.dst, dst, k-1, prev+it.price, out)
	}
}
