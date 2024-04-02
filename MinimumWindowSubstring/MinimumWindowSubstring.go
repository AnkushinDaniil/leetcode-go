package minimumWindowSubstring

func minWindow(s string, t string) string {
	sLen := len(s)

	table := make([]int32, 128)

	for i := range t {
		table[t[i]]++
	}

	counter := len(t)

	resLen := sLen + 1
	resL := 0
	l := 0

	for r := range s {
		if table[s[r]]--; table[s[r]] >= 0 {
			counter--
		}
		for ; counter == 0; l++ {
			if length := r - l + 1; length < resLen {
				resLen = length
				resL = l
			}
			if table[s[l]]++; table[s[l]] > 0 {
				counter++
			}
		}
	}

	if resLen == sLen+1 {
		return ""
	}

	return s[resL : resL+resLen]
}
