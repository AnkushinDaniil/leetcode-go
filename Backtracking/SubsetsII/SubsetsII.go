package subsetsII

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0, 1<<len(nums))
	subset := make([]int, 0, len(nums))

	sort.Ints(nums)

	var backtrack func(int)
	backtrack = func(i int) {
		res = append(res, make([]int, len(subset)))
		copy(res[len(res)-1], subset)
		for j := i; j < len(nums); j++ {
			if j > i && nums[j] == nums[j-1] {
				continue
			}
			subset = append(subset, nums[j])
			backtrack(j + 1)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0)
	return res
}
