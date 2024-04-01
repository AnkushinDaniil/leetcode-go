package minimumWindowSubstring

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)

	if sLen < tLen {
		return ""
	}

	sTable := [26 * 2]uint16{}
	tTable := [26 * 2]uint16{}

	for i := range t {
		tTable[charToInt(t[i])]++
	}

	l := 0
	r := 0

	for r = range s[:tLen] {
		sTable[charToInt(s[r])]++
	}

	res := ""
	resLen := sLen + 1

	for r < sLen {
		isIncluded := true
		for j := 0; j < 26*2 && isIncluded; j++ {
			isIncluded = sTable[j] >= tTable[j]
		}
		if isIncluded {
			if temp := r - l + 1; temp < resLen {
				resLen = temp
				res = s[l : r+1]
			}
			sTable[charToInt(s[l])]--
			l++
		} else if r+1 < sLen {
			r++
			sTable[charToInt(s[r])]++
		} else {
			break
		}
	}

	return res
}

func charToInt(ch uint8) uint8 {
	if ch >= 'a' {
		return ch - 'a' + 26
	} else {
		return ch - 'A'
	}
}
