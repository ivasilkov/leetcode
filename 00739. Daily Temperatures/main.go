package _0739__Daily_Temperatures

func dailyTemperatures(t []int) []int {
	var st []int
	ans := make([]int, len(t))

	for i := len(t) - 1; i >= 0; i-- {
		for len(st) > 0 && t[st[len(st)-1]] <= t[i] {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans[i] = st[len(st)-1] - i
		}
		st = append(st, i)
	}

	return ans
}
