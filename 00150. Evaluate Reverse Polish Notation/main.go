package _0150__Evaluate_Reverse_Polish_Notation

import "strconv"

func evalRPN(tokens []string) int {
	var st []int
	m := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for i := range tokens {
		switch tokens[i] {
		case "+", "*", "/", "-":
			st[len(st)-2] = m[tokens[i]](st[len(st)-2], st[len(st)-1])
			st = st[:len(st)-1]
		default:
			v, _ := strconv.Atoi(tokens[i])
			st = append(st, v)
		}
	}

	return st[0]
}
