package nQueens

import "strings"

func solveNQueens(n int) [][]string {
	queens := make([][2]int, 0, n-1)
	res := make([][]string, 0, n)

	cols := make([]bool, n)
	sum := make([]bool, 2*n)
	dif := make([]bool, 2*n)

	table := make([]string, n)
	row := make([]byte, n)
	for i := 0; i < n; i++ {
		row[i] = '.'
	}
	for i := 0; i < n; i++ {
		row[i] = 'Q'
		table[i] = string(row)
		row[i] = '.'
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, make([]string, n))
			for i := 0; i < n; i++ {
				res[len(res)-1][queens[i][1]] = table[queens[i][0]]
			}
		}

		for c := 0; c < n; c++ {
			if !cols[c] && !sum[row+c] && !dif[n-1+row-c] {
				queens = append(queens, [2]int{c, row})
				cols[c] = true
				sum[row+c] = true
				dif[n-1+row-c] = true
				backtrack(row + 1)
				queens = queens[:len(queens)-1]
				cols[c] = false
				sum[row+c] = false
				dif[n-1+row-c] = false
			}
		}
	}
	backtrack(0)
	return res
}

func solveNQueens2d(n int) [][]string {
	res := make([][]string, 0, n)
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		t := make([]byte, n)
		for j := 0; j < n; j++ {
			t[j] = '.'
		}
		board[i] = t
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, make([]string, n))
			for i := 0; i < n; i++ {
				res[len(res)-1][i] = string(board[i])
			}
		}

	LOOP:
		for col := 0; col < n; col++ {
			for r := 0; r < row; r++ {
				if board[r][col] == 'Q' {
					continue LOOP
				}
			}
			for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
				if board[r][c] == 'Q' {
					continue LOOP
				}
			}
			for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
				if board[r][c] == 'Q' {
					continue LOOP
				}
			}
			board[row][col] = 'Q'
			backtrack(row + 1)
			board[row][col] = '.'
		}
	}

	backtrack(0)
	return res
}

func solveNQueens1d(n int) [][]string {
	res := make([][]string, 0, n)
	board := make([]byte, n*n)
	for i := 0; i < n*n; i++ {
		board[i] = '.'
	}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			res = append(res, make([]string, n))
			for i := 0; i < n; i++ {
				res[len(res)-1][i] = string(board[i : i+n])
			}
		}

	LOOP:
		for col := 0; col < n; col++ {
			for r := 0; r < row; r++ {
				if board[r*n+col] == 'Q' {
					continue LOOP
				}
			}
			for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
				if board[r*n+c] == 'Q' {
					continue LOOP
				}
			}
			for r, c := row-1, col+1; r >= 0 && c < n; r, c = r-1, c+1 {
				if board[r*n+c] == 'Q' {
					continue LOOP
				}
			}
			board[row*n+col] = 'Q'
			backtrack(row + 1)
			board[row*n+col] = '.'
		}
	}

	backtrack(0)
	return res
}

func solveNQueensNeetcode(n int) [][]string {
	ans, curr := make([][]string, 0), make([]string, 0)
	column, diag1, diag2 := make(map[int]int), make(map[int]int), make(map[int]int)
	var backtrack func(y int)
	backtrack = func(y int) {
		if y == n {
			ans = append(ans, append([]string{}, curr...))
		}
		for x := 0; x < n; x++ {
			if column[x] > 0 || diag1[x+y] > 0 || diag2[n+x-y-1] > 0 {
				continue
			}
			column[x], diag1[x+y], diag2[n+x-y-1] = 1, 1, 1
			s := ""
			for i := 0; i < n; i++ {
				if i == x {
					s += "Q"
				} else {
					s += "."
				}
			}
			curr = append(curr, s)
			backtrack(y + 1)
			column[x], diag1[x+y], diag2[n+x-y-1] = 0, 0, 0
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0)
	return ans
}

func solveNQueensLeetcode(n int) [][]string {
	var solutions [][]string
	var backtrack func(row int, board []string)
	backtrack = func(row int, board []string) {
		if row == n {
			var solution []string = make([]string, n, n)
			copy(solution, board)
			solutions = append(solutions, solution)
			return
		}

		board = append(board, strings.Repeat(".", n))

		for col := 0; col < n; col++ { // Try to place the Queen in each column of this row
			// Confirm that queen can't be attacked;
			for prevRow := 1; row-prevRow >= 0; prevRow++ { // Begin check each previous row, starting with row-1
				// need to validate 3 cases for each row:
				// 1. There is no queen in the same column
				// 2. There is no queen in the upper-left diagonal (col-prevRow)
				// 3. There is no queen in the upper-right diagonal (col+prevRow)
				if board[row-prevRow][col] == 81 { // Confirm that the same column is empty
					goto skip // the queen can be attacked, don't place a queen in this column. goto skip.
				}
				if col-prevRow >= 0 &&
					board[row-prevRow][col-prevRow] == 81 { // confirm that the upper-left diagonal is empty
					goto skip // the queen can be attacked, don't place a queen in this column. goto skip.
				}
				if col+prevRow < n &&
					board[row-prevRow][col+prevRow] == 81 { // confirm that the upper-right diagonal is empty
					goto skip // the queen can be attacked, don't place a queen in this column. goto skip.
				}
			}
			// If the queen can't be attacked, the loop will end here and backtracking will continue
			board[row] = board[row][:col] + "Q" + board[row][col+1:] // place a queen in the current col
			backtrack(
				row+1,
				board,
			) // try to place a queen on the next row
			board[row] = board[row][:col] + "." + board[row][col+1:] // remove the queen from this col and continue to next col

			// If the queen can be attacked, the loop will end here and the next column will be investigated
		skip: // this label is used to skip adding a queen if
		}
	}

	backtrack(0, make([]string, 0, n))
	return solutions
}

// Benchmark-10                        3308            358937 ns/op           83472 B/op        375 allocs/op
// Benchmark2d-10                       940           1266957 ns/op          133780 B/op       3537 allocs/op
// BenchmarkNeetcode-10                 274           4288334 ns/op          488297 B/op      67518 allocs/op
// BenchmarkLeetcode-10                 620           1916975 ns/op          480186 B/op      25191 allocs/op
