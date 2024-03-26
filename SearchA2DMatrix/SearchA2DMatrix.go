package binarySearch

func searchMatrix_(matrix [][]int, target int) bool {
	width := len(matrix[0])
	l, r := 0, len(matrix)*width-1
	for l <= r {
		i := (l + r) / 2
		column := i % width
		row := i / width
		if matrix[row][column] == target {
			return true
		} else if matrix[row][column] < target {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	for l <= r {
		i := (l + r) / 2
		if matrix[i][0] == target {
			return true
		} else if matrix[i][0] < target {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	i := max(r, 0)
	l, r = 0, len(matrix[i])-1
	for l <= r {
		j := (l + r) / 2
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = j + 1
		} else {
			r = j - 1
		}
	}
	return false
}
