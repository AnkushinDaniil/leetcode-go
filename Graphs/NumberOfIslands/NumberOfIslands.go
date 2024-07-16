package numberOfIslands

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	queue := make([][2]int, 0, m*n)
	islands := 0

	updateIfOne := func(i, j int) {
		if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == '1' {
			queue = append(queue, [2]int{i, j})
			grid[i][j] = '2'
		}
	}

	checkAndUpdate := func(i, j int) {
		updateIfOne(i+1, j)
		updateIfOne(i, j+1)
		updateIfOne(i-1, j)
		updateIfOne(i, j-1)
	}

	bfs := func(i, j int) {
		grid[i][j] = '2'
		queue = append(queue, [2]int{i, j})
		for len(queue) > 0 {
			i, j = queue[0][0], queue[0][1]
			queue = queue[1:]
			checkAndUpdate(i, j)
		}
	}

	for row := range m {
		for col := range n {
			if grid[row][col] == '1' {
				bfs(row, col)
				islands++
			}
		}
	}

	return islands
}
