package _0658__Find_K_Closest_Elements

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	if arr[0] > x {
		return arr[:k]
	}
	if arr[len(arr)-1] < x {
		return arr[len(arr)-k:]
	}

	l, r := 0, len(arr)-k
	for l < r {
		m := (l + r) / 2
		if x-arr[m] > arr[m+k]-x {
			l = m + 1
		} else {
			r = m
		}
	}
	return arr[l : l+k]
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
