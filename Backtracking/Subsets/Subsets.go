package subsets

import (
	"slices"
)

func subsets(nums []int) [][]int {
  res := make([][]int, 0, 1 << len(nums))
  subset := make([]int, 0, len(nums))

  var dfs func(int)
  dfs = func(start int) {
    res = append(res, slices.Clone(subset))
    for i := start; i < len(nums); i++ {
      subset = append(subset, nums[i])
      dfs(i+1)
      subset = subset[:len(subset)-1]
    }
  }  
  dfs(0)
  return res
}
