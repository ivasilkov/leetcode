package main

func main() {
	numJewelsInStones1("aA", "aAAbbbb")
}

func numJewelsInStones(jewels string, stones string) int {
	m := map[byte]struct{}{}
	for i := 0; i < len(jewels); i++ {
		m[jewels[i]] = struct{}{}
	}

	out := 0
	for i := 0; i < len(stones); i++ {
		if _, ok := m[stones[i]]; ok {
			out++
		}
	}
	return out
}

// 96% memory
func numJewelsInStones1(jewels string, stones string) int {
	arr := make([]bool, 'z'-'A')
	for i := 0; i < len(jewels); i++ {
		arr[jewels[i]-'A'] = true
	}

	out := 0
	for i := 0; i < len(stones); i++ {
		if ok := arr[stones[i]-'A']; ok {
			out++
		}
	}
	return out
}
