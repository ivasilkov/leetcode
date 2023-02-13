package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("1", "1"))
}

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	out := make([]byte, len(a))
	carry := uint8(0)

	for i := 0; i < len(a); i++ {
		sum := carry
		aIdx := len(a) - 1 - i
		bIdx := len(b) - 1 - i

		if len(b)-1 >= i {
			sum += b[bIdx] - '0'
		}
		sum += a[aIdx] - '0'

		if sum >= 2 {
			carry, sum = 1, sum-2
		} else {
			carry = 0
		}

		if sum == 1 {
			out[aIdx] = '1'
		} else {
			out[aIdx] = '0'
		}
	}

	if carry == 1 {
		out = append([]byte{'1'}, out...)
	}
	return string(out)
}

func addBinary1(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	out := make([]byte, len(a))

	carry := false
	for i := 0; i < len(a); i++ {
		aIdx := len(a) - 1 - i
		if len(b)-1 < i {
			out[aIdx], carry = plus(a[aIdx], '0', carry)
			continue
		}

		bIdx := len(b) - 1 - i
		out[aIdx], carry = plus(a[aIdx], b[bIdx], carry)
	}

	if carry {
		out = append([]byte{'1'}, out...)
	}

	return string(out)
}

func plus(a, b byte, carry bool) (byte, bool) {
	switch {
	case a == '1' && b == '1':
		if carry {
			return '1', true
		} else {
			return '0', true
		}
	case a == '1' || b == '1':
		if carry {
			return '0', true
		} else {
			return '1', false
		}
	default:
		if carry {
			return '1', false
		} else {
			return '0', false
		}
	}
}
