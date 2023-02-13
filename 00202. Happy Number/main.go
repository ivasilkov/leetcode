package _0202__Happy_Number

func isHappy(n int) bool {
	m := make(map[int]struct{})

	for {
		tmp := 0
		for n != 0 {
			mod := n % 10
			tmp += mod * mod
			n = n / 10
		}
		n = tmp
		if n == 1 {
			return true
		}
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
	}
}
