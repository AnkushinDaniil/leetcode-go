package stringtointeger

func myAtoi(s string) int {
	i := 0
	for ; i < len(s) && s[i] == ' '; i++ {
	}

	if i == len(s) {
		return 0
	}

	sign := 1
	switch s[i] {
	case '-':
		sign = -1
		i++
	case '+':
		i++
	}

	for ; i < len(s) && s[i] == '0'; i++ {
	}

	res := 0
  bound := 1<<31

  if sign == 1 {
    bound--
  }

	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res *= 10
			res += int(s[i] - '0')
      if res > bound {
        return sign * bound
      }
		} else {
      break
    }
	}

  return sign*res
}
