package verifyinganaliendictionary

func isAlienSorted(words []string, order string) bool {
	alphabet := [256]int{}

	for i := range order {
		alphabet[order[i]] = i
	}

	for i := range words[:len(words)-1] {
		for j := range words[i] {
			if j == len(words[i+1]) {
				return false
			}
			if alphabet[words[i][j]] < alphabet[words[i+1][j]] {
				break
			} else if words[i][j] == words[i+1][j] {
				continue
			}
			return false
		}
	}

	return true
}
