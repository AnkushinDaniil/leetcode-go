package pacificAtlanticWaterFlow

func pacificAtlantic(heights [][]int) [][]int {
	rows, cols := len(heights), len(heights[0])
	isPacific := make([][]bool, rows)
	isAtlantic := make([][]bool, rows)
	res := make([][]int, 0, rows+cols)

	for row := 0; row < rows; row++ {
		isPacific[row] = make([]bool, cols)
		isAtlantic[row] = make([]bool, cols)
	}

	var dfs func(int, int)
	dfs = func(row, col int) {
		isPacific[row][col] = true
		if row+1 < rows && !isPacific[row+1][col] && heights[row+1][col] >= heights[row][col] {
			dfs(row+1, col)
		}
		if col+1 < cols && !isPacific[row][col+1] && heights[row][col+1] >= heights[row][col] {
			dfs(row, col+1)
		}
		if 0 <= row-1 && !isPacific[row-1][col] && heights[row-1][col] >= heights[row][col] {
			dfs(row-1, col)
		}
		if 0 <= col-1 && !isPacific[row][col-1] && heights[row][col-1] >= heights[row][col] {
			dfs(row, col-1)
		}
	}

	for row := 0; row < rows; row++ {
		if !isPacific[row][0] {
			dfs(row, 0)
		}
	}
	for col := 1; col < cols; col++ {
		if !isPacific[0][col] {
			dfs(0, col)
		}
	}

	dfs = func(row, col int) {
		isAtlantic[row][col] = true
		if isPacific[row][col] {
			res = append(res, []int{row, col})
		}
		if row+1 < rows && !isAtlantic[row+1][col] && heights[row+1][col] >= heights[row][col] {
			dfs(row+1, col)
		}
		if col+1 < cols && !isAtlantic[row][col+1] && heights[row][col+1] >= heights[row][col] {
			dfs(row, col+1)
		}
		if 0 <= row-1 && !isAtlantic[row-1][col] && heights[row-1][col] >= heights[row][col] {
			dfs(row-1, col)
		}
		if 0 <= col-1 && !isAtlantic[row][col-1] && heights[row][col-1] >= heights[row][col] {
			dfs(row, col-1)
		}
	}

	for row := 0; row < rows; row++ {
		if !isAtlantic[row][cols-1] {
			dfs(row, cols-1)
		}
	}
	for col := 0; col < cols-1; col++ {
		if !isAtlantic[rows-1][col] {
			dfs(rows-1, col)
		}
	}

	return res
}
