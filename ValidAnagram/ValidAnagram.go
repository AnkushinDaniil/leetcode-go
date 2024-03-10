package validAnagram

func isAnagram(s string, t string) bool {
	m := len(s)
	n := len(t)

	if n != m {
		return false
	}

	table := make([]int, 26)

	for i := 0; i < m; i++ {
		table[s[i]-'a']++
		table[t[i]-'a']--
	}

	for _, v := range table {
		if v != 0 {
			return false
		}
	}

	return true
}
