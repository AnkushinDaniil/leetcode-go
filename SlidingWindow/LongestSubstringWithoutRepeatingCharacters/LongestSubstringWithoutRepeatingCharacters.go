package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n < 2 {
		return n
	}

	res, l := 1, 0

	for r := 1; r < n; r++ {
		for i := l; i < r; i++ {
			if s[i] == s[r] {
				if t := r - l; t > res {
					res = t
				}
				l = i + 1
				break
			}
		}
	}

	if t := n - l; t > res {
		res = t
	}

	return res
}
