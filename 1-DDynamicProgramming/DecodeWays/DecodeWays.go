package decodeWays

func numDecodings(s string) int {
	a := 1
	b := 0

	if s[0] != '0' {
		b = 1
	}

	for i := 2; i <= len(s); i++ {
		if !(s[i-2] == '1' || s[i-2] == '2' && s[i-1] < '7') {
			a = 0
		}

		if s[i-1] != '0' {
			a += b
		}
		a, b = b, a
	}

	return b
}

// ----------------------------------

func numDecodingsMy(s string) int {
	if len(s) == 1 {
		if s[0] == '0' {
			return 0
		} else {
			return 1
		}
	}
	dp1, dp2, dp3 := 0, 0, 1
	if s[len(s)-1] == '0' {
		dp3 = 0
	}
	if s[len(s)-2] == '0' {
		dp2 = 0
	} else {
		dp2 = dp3
	}
	if s[len(s)-2] == '1' || s[len(s)-2] == '2' && s[len(s)-1] >= '0' && s[len(s)-1] <= '6' {
		dp2 = dp3 + 1
	}

	if len(s) == 2 {
		return dp2
	}

	for i := len(s) - 3; i >= 0; i-- {
		if s[i] == '0' {
			dp1 = 0
		} else {
			dp1 = dp2
		}

		if s[i] == '1' || s[i] == '2' && s[i+1] >= '0' && s[i+1] <= '6' {
			dp1 += dp3
		}
		dp2, dp3 = dp1, dp2
	}

	return dp1
}

// ----------------------------------

func numDecodingsLeetcodeTime(s string) int {
	return dfs(s, 0, make(map[int]int))
}

func dfs(s string, idx int, memo map[int]int) int {
	l := len(s)
	if idx == l {
		return 1
	} else if v, ok := memo[idx]; ok {
		return v
	}

	cnt := 0
	for i := idx + 1; i <= min(idx+2, l); i++ {
		if !isValidEncoding(s[idx:i]) {
			continue
		}

		cnt += dfs(s, i, memo)
	}

	memo[idx] = cnt
	return cnt
}

func isValidEncoding(s string) bool {
	if len(s) == 1 {
		return s[0] >= byte('1') && s[0] <= byte('9')
	} else if len(s) == 2 {
		if s[0] == '1' {
			return s[1] >= byte('0') && s[1] <= byte('9')
		} else if s[0] == '2' {
			return s[1] >= byte('0') && s[1] <= byte('6')
		} else {
			return false
		}
	} else {
		return false
	}
}

// ----------------------------------

func numDecodingsLeetcodeMemory(s string) int {
	res := []int{1, 0}
	if s[0] != 48 {
		res[1] = 1
	}

	for i := 2; i <= len(s); i++ {
		if !(s[i-2] == 49 || s[i-2] == 50 && s[i-1] < 55) {
			res[i%2] = 0
		}
		if s[i-1] != 48 {
			res[i%2] += res[(i+1)%2]
		}
	}

	return res[len(s)%2]
}
