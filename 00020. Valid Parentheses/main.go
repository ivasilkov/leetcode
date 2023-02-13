package _0020__Valid_Parentheses

func isValid(s string) bool {
	var stack []byte

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if s[i] == ')' && last != '(' {
				return false
			}
			if s[i] == ']' && last != '[' {
				return false
			}
			if s[i] == '}' && last != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func isValid1(s string) bool {
	var stack []byte
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
