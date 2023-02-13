package main

import "sort"

func main() {

}

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i := 0; i < len(capacity); i++ {
		capacity[i] -= rocks[i]
	}
	sort.Ints(capacity)

	for i := 0; i < len(capacity); i++ {
		if capacity[i] > additionalRocks {
			return i
		}
		additionalRocks -= capacity[i]
	}
	return len(capacity)
}
