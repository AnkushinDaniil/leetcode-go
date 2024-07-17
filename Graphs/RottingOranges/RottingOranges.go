package rottingOranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	time, counter := 0, 0
	queue := make([][2]int, 0, 4*max(rows, cols))

	for row := range rows {
		for col := range cols {
			if grid[row][col] == 2 {
				queue = append(queue, [2]int{row, col})
			} else if grid[row][col] == 1 {
				counter++
			}
		}
	}

	if counter == 0 {
		return 0
	}

	checkAndUpdate := func(i, j int) {
		if i >= 0 && i < rows && j >= 0 && j < cols &&
			grid[i][j] == 1 {
			queue = append(queue, [2]int{i, j})
			grid[i][j] = 2
			counter--
		}
	}

	checkAndUpdateNeighbors := func(i, j int) {
		checkAndUpdate(i+1, j)
		checkAndUpdate(i, j+1)
		checkAndUpdate(i-1, j)
		checkAndUpdate(i, j-1)
	}

	for len(queue) > 0 {
		lenQueue := len(queue)
		time++
		for i := 0; i < lenQueue; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			checkAndUpdateNeighbors(row, col)
		}
	}

	if counter == 0 {
		return time - 1
	} else {
		return -1
	}
}
