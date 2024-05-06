package combinationSum

func permute(nums []int) [][]int {
	res := make([][]int, 0, factorial(len(nums)))

	var backtrack func(int)
	backtrack = func(i int) {
		if i == len(nums) {
			res = append(res, make([]int, len(nums)))
			copy(res[len(res)-1], nums)
			return
		}
		for j := i; j < len(nums); j++ {
			nums[j], nums[i] = nums[i], nums[j]
			backtrack(i + 1)
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	backtrack(0)
	return res
}

func factorial(n int) int {
	res := 1
	for ; n > 1; n-- {
		res *= n
	}
	return res
}
