package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	top := 0
	bottom := len(matrix)
	left := 0
	right := len(matrix[0])

	res := make([]int, 0, bottom*right)

	bottom--
	right--

	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		if top++; top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		if right--; right < left {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		if bottom--; bottom < top {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		if left++; left > right {
			break
		}
	}

	return res
}
