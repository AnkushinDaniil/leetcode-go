package climbingStairs

func climbStairsDFS(n int) int {
	res := 0
	table := make([]int, n+1)
	isCalculated := make([]bool, n+1)

	var dfs func()
	dfs = func() {
		if n < 0 {
			return
		}
		if isCalculated[n] {
			res += table[n]
			return
		} else {
			if 0 == n {
				res++
			} else if 0 < n {
				n--
				dfs()
				n--
				dfs()
				n += 2
			}
			r := res
			table[n] = r
			isCalculated[n] = true
		}
	}

	dfs()

	return res
}

func climbStairs(n int) int {
	step1, step2 := 0, 1
	for ; n > 0; n-- {
		step1, step2 = step2, step1+step2
	}
	return step2
}
