package main

import "math"

func main() {

}

// so fast
func reverse(x int) int {
	mult := 1
	if x < 0 {
		x = -x
		mult = -1
	}

	out := 0
	for x != 0 {
		out = out*10 + x%10
		x = x / 10
	}

	if out > math.MaxInt32 || out < math.MinInt32 {
		return 0
	}
	return out * mult
}

func reverse2(x int) int {
	out := 0
	for x != 0 {
		out = out*10 + x%10
		x = x / 10
	}
	if out > math.MaxInt32 || out < math.MinInt32 {
		return 0
	}
	return out
}
