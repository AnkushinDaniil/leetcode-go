package longestIncreasingSubsequence

func lengthOfLIS(nums []int) int {
	memo := make([]int, len(nums))
	memo[len(nums)-1] = 1
	res := 1
	for i := len(nums) - 2; i >= 0; i-- {
		maximum := 0
		for j := i + 1; j < len(memo); j++ {
			if nums[i] < nums[j] && memo[j] > maximum {
				maximum = memo[j]
			}
		}
		memo[i] = maximum + 1
		if memo[i] > res {
			res = memo[i]
		}
	}
	return res
}
