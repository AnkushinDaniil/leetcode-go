package generateParentheses

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	stack := make([]uint8, 0, 2*n)

	var calc func(int, int)
	calc = func(o, c int) {
		if o == n && c == n {
			res = append(res, string(stack))
			return
		}
		if o < n {
			stack = append(stack, '(')
			calc(o+1, c)
			stack = stack[:len(stack)-1]
		}
		if c < o {
			stack = append(stack, ')')
			calc(o, c+1)
			stack = stack[:len(stack)-1]
		}
	}

	calc(0, 0)
	return res
}
