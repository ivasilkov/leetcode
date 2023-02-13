package main

import (
	"container/heap"
	"sort"
)

func main() {
	arr := [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}
	_ = getOrder(arr)
}

func getOrder(tasks [][]int) []int {
	list := make([][3]int, len(tasks))
	for i, v := range tasks {
		list[i] = [3]int{v[0], v[1], i}
	}
	sort.Slice(list, func(i int, j int) bool {
		return list[i][0] < list[j][0]
	})

	h := &minheap{}
	cur := list[0][0]
	i := 0
	var res []int
	for h.Len() > 0 || i < len(list) {
		for i < len(list) && list[i][0] <= cur {
			heap.Push(h, list[i])
			i++
		}
		if h.Len() == 0 && i < len(list) {
			cur = list[i][0]
			continue
		}
		t := heap.Pop(h).([3]int)
		cur += t[1]
		res = append(res, t[2])
	}
	return res
}

type minheap [][3]int

func (h minheap) Len() int {
	return len(h)
}

func (h minheap) Less(i int, j int) bool {
	if h[i][1] == h[j][1] {
		return h[i][2] < h[j][2]
	}
	return h[i][1] < h[j][1]
}

func (h minheap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minheap) Push(a interface{}) {
	*h = append(*h, a.([3]int))
}

func (h *minheap) Pop() interface{} {
	l := len(*h)
	res := (*h)[l-1]
	*h = (*h)[:l-1]
	return res
}
