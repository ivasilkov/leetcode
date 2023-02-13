package main

import "sort"

func main() {
	_ = answerQueries([]int{4, 5, 2, 1}, []int{3, 10, 21})
}

func answerQueries(nums []int, queries []int) []int {
	n, m := len(nums), len(queries)

	sort.Ints(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}

	out := make([]int, m)
	for i := 0; i < m; i++ {
		out[i] = sort.Search(len(nums), func(j int) bool {
			return nums[j] > queries[i]
		})
	}
	return out
}
