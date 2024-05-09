package wordSearch

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	w := 0
	res := false

	// added part
	start, end := 1, 0

	for ; start < len(word) && word[start-1] == word[start]; start++ {
	}
	for ; end < len(word)-1 && word[len(word)-1-end] == word[len(word)-2-end]; end++ {
	}

	if start > end {
		wordBytes := []byte(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			wordBytes[i], wordBytes[j] = wordBytes[j], wordBytes[i]
		}
		word = string(wordBytes)
	}
	// added part

	var req func(int, int)
	req = func(i, j int) {
		if res || i < 0 || i >= n || j < 0 || j >= m ||
			word[w] != board[i][j] || board[i][j] == 0 {
			return
		}
		w++
		if w == len(word) {
			res = true
			return
		}
		ch := board[i][j]
		board[i][j] = 0
		req(i+1, j)
		req(i-1, j)
		req(i, j+1)
		req(i, j-1)
		w--
		board[i][j] = ch
	}

	for x := 0; !res && x < n; x++ {
		for y := 0; !res && y < m; y++ {
			req(x, y)
		}
	}
	return res
}

func existSimple(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	w := 0
	res := false

	var req func(int, int)
	req = func(i, j int) {
		if res || i < 0 || i >= n || j < 0 || j >= m ||
			word[w] != board[i][j] || board[i][j] == 0 {
			return
		}
		w++
		if w == len(word) {
			res = true
			return
		}
		ch := board[i][j]
		board[i][j] = 0
		req(i+1, j)
		req(i-1, j)
		req(i, j+1)
		req(i, j-1)
		w--
		board[i][j] = ch
	}

	for x := 0; !res && x < n; x++ {
		for y := 0; !res && y < m; y++ {
			req(x, y)
		}
	}
	return res
}
