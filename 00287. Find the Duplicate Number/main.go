package main

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 0, n-1 // идем не по входному массиву, а по 1..n
	for l < r {
		m := (l + r) / 2
		cnt := 0
		for _, v := range nums {
			if v <= m {
				cnt++
			}
		}
		if cnt <= m {
			l = m + 1 // значит для до m нет дубликатов
		} else {
			r = m // значит на этом интервале были не уникальные
		}
	}
	return l
}

// if only 1 duplicate
func findDuplicate1(nums []int) int {
	expected, sum := 0, 0
	for i, v := range nums {
		expected += i + 1
		sum += v
	}
	return len(nums) - (expected - sum)
}
