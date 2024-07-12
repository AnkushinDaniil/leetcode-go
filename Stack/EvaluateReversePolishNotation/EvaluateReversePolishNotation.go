package evaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for i := range tokens {
		switch tokens[i] {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "-":
			stack[len(stack)-2] -= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "/":
			stack[len(stack)-2] /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			res, _ := strconv.Atoi(tokens[i])
			stack = append(stack, res)
		}
	}
	return stack[0]
}
