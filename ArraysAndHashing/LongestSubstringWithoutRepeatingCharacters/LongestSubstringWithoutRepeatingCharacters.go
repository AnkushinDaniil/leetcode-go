package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	res := 0
	count := [256]uint16{}
	r, l := 0, 0
	for r < len(s) {
		count[s[r]]++
		for count[s[r]] > uint16(1) {
			count[s[l]]--
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}
	return res
}
