package _0907__Sum_of_Subarray_Minimums

type stack []int

func (s *stack) push(v int) {
	*s = append(*s, v)
}
func (s *stack) pop() int {
	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return v
}
func (s *stack) peak() int {
	v := (*s)[len(*s)-1]
	return v
}

func sumSubarrayMins(arr []int) int {
	out := 0
	var st stack
	st.push(-1)

	for i := range arr {
		for len(st) > 1 && arr[i] <= arr[st.peak()] {
			idx := st.pop()
			out += (i - idx) * (idx - st.peak()) * arr[idx]
		}
		st = append(st, i)
	}

	for len(st) > 1 {
		idx := st.pop()
		out += (len(arr) - idx) * (idx - st.peak()) * arr[idx]
	}

	return out % 1000000007
}
