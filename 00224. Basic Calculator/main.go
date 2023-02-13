package _0224__Basic_Calculator

func calculate(s string) int {
	num, sign := 0, 1
	stack := []int{0}

	for i := range s {
		switch s[i] {
		case ' ':
			continue
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num*10 + int(s[i]-'0')
		case '0':
			num = num * 10
		case '+':
			stack[len(stack)-1] += num * sign
			sign, num = 1, 0
		case '-':
			stack[len(stack)-1] += num * sign
			sign, num = -1, 0
		case '(':
			stack = append(stack, sign, 0)
			sign, num = 1, 0
		case ')':
			v := (stack[len(stack)-1] + num*sign) * stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack[len(stack)-1] += v
			sign, num = 1, 0
		}
	}

	return stack[len(stack)-1] + num*sign
}
