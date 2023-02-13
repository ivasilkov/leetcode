package _0446__Arithmetic_Slices_II___Subsequence

func main() {

}

func numberOfArithmeticSlices(nums []int) int {
	m := make([]map[int]int, len(nums))
	for i := range m {
		m[i] = map[int]int{}
	}

	out := 0
	for i, a := range nums {
		for j, b := range nums[:i] {
			m[i][a-b] += m[j][a-b] + 1
			out += m[j][a-b]

			//diff := nums[i] - nums[j]
			//a := m[j][diff]
			//b := m[i][diff]
			//out += a
			//m[i][diff] = a + b + 1
		}
	}
	return out
}
