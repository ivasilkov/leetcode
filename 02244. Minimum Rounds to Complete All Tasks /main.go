package main

func main() {

}

func minimumRounds(tasks []int) int {
	m := map[int]int{}
	for _, t := range tasks {
		m[t]++
	}

	out := 0
	for _, v := range m {
		if v == 1 {
			return -1
		}
		out += v / 3
		if v%3 != 0 {
			out++
		}
	}
	return out
}
