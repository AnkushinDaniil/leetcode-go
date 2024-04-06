package evaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	n := -1
	stack := make([]int, len(tokens))

	for _, token := range tokens {
		switch token {
		case "+":
			stack[n-1] = stack[n-1] + stack[n]
			n--
		case "-":
			stack[n-1] = stack[n-1] - stack[n]
			n--
		case "*":
			stack[n-1] = stack[n-1] * stack[n]
			n--
		case "/":
			stack[n-1] = stack[n-1] / stack[n]
			n--
		default:
			n++
			stack[n], _ = strconv.Atoi(token)
		}
	}

	return stack[n]
}
