package subsets

func subsets(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	subset := make([]int, 0, len(nums))
	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			res = append(res, make([]int, len(subset)))
			copy(res[len(res)-1], subset)
			return
		}
		subset = append(subset, nums[i])
		dfs(i + 1)
		subset = subset[:len(subset)-1]
		dfs(i + 1)
	}
	dfs(0)
	return res
}
