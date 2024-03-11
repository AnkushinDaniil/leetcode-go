package groupAnagrams

func groupAnagrams(strs []string) [][]string {
	table := make(map[string][]string, len(strs))

	for _, s := range strs {
		alphabet := make([]byte, 26)

		for i := range s {
			alphabet[s[i]-'a']++
		}

		table[string(alphabet)] = append(table[string(alphabet)], s)
	}

	res := make([][]string, 0, len(table))

	for _, v := range table {
		res = append(res, v)
	}

	return res
}
