package validParenthesisString

func checkValidString(s string) bool {
	counter1, counter2 := 0, 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			counter1++
			counter2++
		case ')':
			counter1--
			counter2--
		default:
			counter1--
			counter2++
		}
		if counter2 < 0 {
			return false
		}
		if counter1 < 0 {
			counter1 = 0
		}
	}
	return counter1 == 0
}
