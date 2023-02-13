package _0744__Find_Smallest_Letter_Greater_Than_Target

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	for l < r {
		m := (l + r) / 2
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}
	if letters[r] > target {
		return letters[r]
	}
	return letters[0]
}
