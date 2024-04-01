package minimumWindowSubstring

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)

	if sLen < tLen {
		return ""
	}

	table := make(map[uint8]int, 0)

	for i := range t {
		table[t[i]]++
	}

	counter := len(table)

	resLen := sLen + 1
	resL := 0

	indexes := make([]int, 0)
	index := 0

	for r := range s {
		if _, ok := table[s[r]]; ok {
			indexes = append(indexes, r)
			if counter > 0 {
				if table[s[r]]--; table[s[r]] == 0 {
					counter--
				}
			}
			for ; counter == 0; index++ {
				l := indexes[index]
				if length := r - l + 1; length < resLen {
					resLen = length
					resL = l
				}
				if table[s[l]]++; table[s[l]] == 1 {
					counter++
				}
			}
		}
	}

	if resLen == sLen+1 {
		return ""
	}

	return s[resL : resL+resLen]
}
