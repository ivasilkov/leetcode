package main

import "sort"

func main() {

}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	out := 0
	for _, cost := range costs {
		if cost > coins {
			break
		}
		coins -= cost
		out++
	}

	return out
}
