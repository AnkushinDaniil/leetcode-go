package fibonaccinumber

func fibN(n int) int {
	if n < 2 {
		return n
	}

	a, b := 0, 1
	for range n - 1 {
		a, b = b, a+b
	}

	return b
}

type mat [2][2]int

func fib(n int) int {
	q := mat{
		{1, 1},
		{1, 0},
	}
	if n == 0 {
		return 0
	}
	return povMat(q, n-1)[0][0]
}

func mulMat(m1, m2 mat) mat {
	return mat{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]},
	}
}

func povMat(m mat, p int) mat {
	res := mat{
		{1, 0},
		{0, 1},
	}
	for ; p > 0; p >>= 1 {
		if p&1 == 1 {
			res = mulMat(res, m)
		}
		m = mulMat(m, m)
	}
	return res
}
