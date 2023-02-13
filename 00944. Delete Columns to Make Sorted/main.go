package main

func main() {

}

func minDeletionSize(strs []string) int {
	if len(strs) <= 1 {
		return 0
	}

	out := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				out++
				break
			}
		}
	}
	return out
}
