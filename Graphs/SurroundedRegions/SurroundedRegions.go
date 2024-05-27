package surroundedRegions

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row < 0 || col < 0 || row >= rows || col >= cols ||
			board[row][col] != 'O' {
			return
		}
		board[row][col] = 'V'
		for i := 0; i < 4; i++ {
			dfs(row+dirs[i][0], col+dirs[i][1])
		}
	}

	for row := 0; row < rows; row++ {
		dfs(row, 0)
		dfs(row, cols-1)
	}
	for col := 1; col < cols-1; col++ {
		dfs(0, col)
		dfs(rows-1, col)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == 'V' {
				board[row][col] = 'O'
			} else {
				board[row][col] = 'X'
			}
		}
	}
}
