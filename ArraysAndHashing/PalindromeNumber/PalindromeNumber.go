package palindromeNumber

func isPalindrome(x int) bool {
	var y int
	tmp := x

	for tmp > 0 {
		y = y*10 + tmp%10
		tmp = tmp / 10
	}

	return x == y
}
