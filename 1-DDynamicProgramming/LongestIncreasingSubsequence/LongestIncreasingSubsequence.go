package longestIncreasingSubsequence

func lengthOfLIS(nums []int) int {
	arr := make([]int, 1, len(nums))
	arr[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		l, r := 0, len(arr)
		for l < r {
			m := (l + r) / 2
			if arr[m] < nums[i] {
				l = m + 1
			} else {
				r = m
			}
		}

		if l == len(arr) {
			arr = append(arr, nums[i])
		} else {
			arr[l] = nums[i]
		}
	}

	return len(arr)
}

func lengthOfLIS_DP(nums []int) int {
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
