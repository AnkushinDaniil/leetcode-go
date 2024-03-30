package longestRepeatingCharacterReplacement

func characterReplacement(s string, k int) int {
	table, res, l, m := make([]int, 26), 0, 0, 0

	for r := range s {
		rch := s[r] - 'A'

		if table[rch]++; table[rch] > m {
			m = table[rch]
		}
		if r-l+1 > m+k {
			table[s[l]-'A']--
			l++
		}
		if d := r - l + 1; d > res {
			res = d
		}
	}

	return res
}
