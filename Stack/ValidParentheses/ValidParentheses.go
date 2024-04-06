package validParentheses

func isValid(s string) bool {
	stack := make([]uint8, len(s))
	n := -1

	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			n++
			stack[n] = s[i]
		case ')', ']', '}':
			var ch uint8
			switch s[i] {
			case ')':
				ch = '('
			case ']':
				ch = '['
			case '}':
				ch = '{'
			}
			if n >= 0 && stack[n] == ch {
				n--
			} else {
				return false
			}
		}
	}

	return n == -1
}
