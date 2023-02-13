package main

import "container/heap"

func main() {
	arr := []int{5, 4, 9}
	k := 2
	_ = minStoneSum(arr, k)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() any {
	last := len(*h) - 1
	out := (*h)[last]
	*h = (*h)[:last]
	return out
}

func minStoneSum(piles []int, k int) int {
	h := IntHeap(piles)
	heap.Init(&h)

	for i := 0; i < k; i++ {
		h[0] -= h[0] / 2
		heap.Fix(&h, 0)
	}

	out := 0
	for _, p := range piles {
		out += p
	}
	return out
}

func minStoneSum1(piles []int, k int) int {
	for i := 0; i < k; i++ {
		max := 0
		for j := range piles {
			if piles[j] > piles[max] {
				max = j
			}
		}
		piles[max] = piles[max]/2 + piles[max]%2
	}

	out := 0
	for _, p := range piles {
		out += p
	}
	return out
}
