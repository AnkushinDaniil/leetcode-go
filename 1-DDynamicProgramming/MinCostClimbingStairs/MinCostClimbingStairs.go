package minCostClimbingStairs

func minCostClimbingStairsDFS(cost []int) int {
	n := len(cost)
	table := make([]int, n)
	isCalculated := make([]bool, n)

	n--
	table[n] = cost[n]
	isCalculated[n] = true
	n--
	table[n] = cost[n]
	isCalculated[n] = true
	n--

	var dfs func()
	dfs = func() {
		if n < 0 {
			return
		}
		if isCalculated[n] {
			return
		} else {
			table[n] = cost[n] + min(table[n+1], table[n+2])
			isCalculated[n] = true
			n--
			dfs()
			n--
			dfs()
			n += 2
		}
	}

	dfs()

	return min(table[0], table[1])
}

func minCostClimbingStairs(cost []int) int {
	for i := len(cost) - 3; i >= 0; i-- {
		if cost[i+1] < cost[i+2] {
			cost[i] += cost[i+1]
		} else {
			cost[i] += cost[i+2]
		}
	}

	if cost[0] < cost[1] {
		return cost[0]
	} else {
		return cost[1]
	}
}
