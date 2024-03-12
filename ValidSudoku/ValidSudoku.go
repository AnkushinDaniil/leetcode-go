package validSudoku

func isValidSudoku(board [][]byte) bool {
	table := make([][]map[byte]struct{}, 3)

	for i := 0; i < 3; i++ {
		table[i] = make([]map[byte]struct{}, 9)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			table[i][j] = make(map[byte]struct{})
		}
	}

	for i, str := range board {
		for j, ch := range str {
			if ch == '.' {
				continue
			}
			k := i/3*3 + j/3

			for k, v := range map[int]int{
				0: i,
				1: j,
				2: k,
			} {
				if _, ok := table[k][v][ch]; ok {
					return false
				} else {
					table[k][v][ch] = struct{}{}
				}
			}
		}
	}

	return true
}
