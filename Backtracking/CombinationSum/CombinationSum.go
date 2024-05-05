package combinationSum

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	subset := make([]int, 0)
	sum := 0

	var dfs func(int)
	dfs = func(i int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, make([]int, len(subset)))
			copy(res[len(res)-1], subset)
			return
		}
		for j := i; j < len(candidates); j++ {
			subset = append(subset, candidates[j])
			sum += candidates[j]
			dfs(j)
			subset = subset[:len(subset)-1]
			sum -= candidates[j]
		}
	}

	dfs(0)

	return res
}
