package main

import (
	"fmt"
	"sort"
)

func main() {
	r := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Println(r)
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	freq := make([][]int, len(nums))
	for n, f := range m {
		freq[f] = append(freq[f], n)
	}

	out := make([]int, 0)
	for i := len(freq) - 1; i > 0; i-- {
		if len(freq[i]) == 0 {
			continue
		}
		out = append(out, freq[i]...)
	}
	return out[:k]
}

func topKFrequent1(nums []int, k int) []int {
	m := map[int]int{}
	arr := make([]int, 0, k)
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			arr = append(arr, n)
		}
		m[n]++
	}

	sort.Slice(arr, func(i, j int) bool {
		return m[arr[i]] > m[arr[j]]
	})

	return arr[:k]
}
