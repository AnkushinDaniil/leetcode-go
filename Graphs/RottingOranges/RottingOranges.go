package rottingOranges

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	time, counter := 0, 0
	queue1 := make([][2]int, 0, 4*max(rows, cols))
	queue2 := make([][2]int, 0, 4*max(rows, cols))

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 2 {
				queue2 = append(queue2, [2]int{row, col})
			} else if grid[row][col] == 1 {
				counter++
			}
		}
	}

	if counter == 0 {
		return 0
	}

	for len(queue1) > 0 || len(queue2) > 0 {
		queue1, queue2 = queue2, queue1
		time++
		for len(queue1) > 0 {
			row, col := queue1[0][0], queue1[0][1]
			queue1 = queue1[1:]
			if row+1 < rows && grid[row+1][col] == 1 {
				queue2 = append(queue2, [2]int{row + 1, col})
				grid[row+1][col] = 2
				counter--
			}
			if col+1 < cols && grid[row][col+1] == 1 {
				queue2 = append(queue2, [2]int{row, col + 1})
				grid[row][col+1] = 2
				counter--
			}
			if row-1 >= 0 && grid[row-1][col] == 1 {
				queue2 = append(queue2, [2]int{row - 1, col})
				grid[row-1][col] = 2
				counter--
			}
			if col-1 >= 0 && grid[row][col-1] == 1 {
				queue2 = append(queue2, [2]int{row, col - 1})
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
