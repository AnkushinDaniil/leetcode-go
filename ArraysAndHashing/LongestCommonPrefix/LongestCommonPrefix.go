package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 0; i < len(strs) && len(res) > 0; i++ {
		res = res[:min(len(res), len(strs[i]))]
		for j := 0; j < len(res); j++ {
			if res[j] != strs[i][j] {
				res = res[:j]
				break
			}
		}
	}
	return res
}
