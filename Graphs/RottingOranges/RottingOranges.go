package rottingOranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	time, counter := 0, 0
	queue := make([][2]int, 0, 4*max(rows, cols))

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
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

	for len(queue) > 0 {
		lenQueue := len(queue)
		time++
		for i := 0; i < lenQueue; i++ {
			row, col := queue[0][0], queue[0][1]
			queue = queue[1:]
			if row+1 < rows && grid[row+1][col] == 1 {
				queue = append(queue, [2]int{row + 1, col})
				grid[row+1][col] = 2
				counter--
			}
			if col+1 < cols && grid[row][col+1] == 1 {
				queue = append(queue, [2]int{row, col + 1})
				grid[row][col+1] = 2
				counter--
			}
			if row-1 >= 0 && grid[row-1][col] == 1 {
				queue = append(queue, [2]int{row - 1, col})
				grid[row-1][col] = 2
				counter--
			}
			if col-1 >= 0 && grid[row][col-1] == 1 {
				queue = append(queue, [2]int{row, col - 1})
				grid[row][col-1] = 2
				counter--
			}
		}
	}
	if counter == 0 {
		return time - 1
	} else {
		return -1
	}
}
