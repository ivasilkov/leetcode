package _0050__Pow_x__n__

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	out := float64(1)
	for n > 0 {
		if n&1 == 1 {
			out *= x
		}
		x *= x
		n = n >> 1
	}
	return out
}
