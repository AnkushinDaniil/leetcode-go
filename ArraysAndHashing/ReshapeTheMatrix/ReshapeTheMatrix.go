package reshapethematrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat) == 0 || len(mat)*len(mat[0]) != r*c {
    return mat
  }
  
  res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
	}

	for i := range mat {
		for j, v := range mat[i] {
			idx := i*len(mat[i]) + j
			res[idx/c][idx%c] = v
		}
	}

	return res
}
