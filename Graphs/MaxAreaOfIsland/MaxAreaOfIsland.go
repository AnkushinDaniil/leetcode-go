package maxAreaOfIsland

func maxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	maxIslandArea := 0
	q := make([][2]int, 0, 4*max(rows, cols))

	bfs := func(i, j int) {
		area := 1
		grid[i][j] = 2
		q = q[:0]
		q = append(q, [2]int{i, j})
		for len(q) > 0 {
			i, j := q[0][0], q[0][1]
			q = q[1:]
			if i+1 < rows && grid[i+1][j] == 1 {
				q = append(q, [2]int{i + 1, j})
				grid[i+1][j] = 2
				area++
			}
			if j+1 < cols && grid[i][j+1] == 1 {
				q = append(q, [2]int{i, j + 1})
				grid[i][j+1] = 2
				area++
			}
			if i-1 >= 0 && grid[i-1][j] == 1 {
				q = append(q, [2]int{i - 1, j})
				grid[i-1][j] = 2
				area++
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				q = append(q, [2]int{i, j - 1})
				grid[i][j-1] = 2
				area++
			}
		}
		if area > maxIslandArea {
			maxIslandArea = area
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				bfs(row, col)
			}
		}
	}

	return maxIslandArea
}
