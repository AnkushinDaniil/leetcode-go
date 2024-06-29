package romanToInteger

func backspaceCompare(s string, t string) bool {
	m, n := len(s)-1, len(t)-1
	i := 0
	for ; s[i] == '#' && i < m; i++ {
	}
	j := 0
	for ; t[j] == '#' && j < n; j++ {
	}
	ctr := 0
	for {
		ctr = 0
		for m >= i && s[m] == '#' {
			ctr++
			m--
			for m >= i && s[m] == '#' {
				ctr++
				m--
			}
			for m >= i && s[m] != '#' && ctr > 0 {
				ctr--
				m--
			}
		}
		ctr = 0
		for n >= j && t[n] == '#' {
			ctr++
			n--
			for n >= j && t[n] == '#' {
				ctr++
				n--
			}
			for n >= j && t[n] != '#' && ctr > 0 {
				ctr--
				n--
			}
		}
		if m == i-1 || n == j-1 {
			return (m == i-1) == (n == j-1)
		}
		if s[m] != t[n] {
			return false
		}
		m--
		n--
	}
}

func backspaceCompareOld(s string, t string) bool {
	strArr := [2]string{s, t}
	for i := 0; i < len(strArr); i++ {
		bytesArr := []byte(strArr[i])
		ctr := 0
		for j := 0; j < len(strArr[i]); j++ {
			if bytesArr[j] == '#' {
				ctr = max(0, ctr-1)
			} else {
				bytesArr[ctr] = bytesArr[j]
				ctr++
			}
		}
		strArr[i] = string(bytesArr[:ctr])
	}
	return strArr[0] == strArr[1]
}
