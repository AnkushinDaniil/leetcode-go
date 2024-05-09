package wordSearch

func exist(board [][]byte, word string) bool {
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
