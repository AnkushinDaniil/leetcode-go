package minimumoperationstomakemedianofarrayequaltok

import "sort"

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	n := len(nums)
	res := 0

	res += nums[n/2] - k
	if res > 0 {
		for i := 1; n/2-i >= 0 && nums[n/2-i] > k; i++ {
			res += nums[n/2-i] - k
		}
	} else {
		res = -res
		for i := 1; n/2+i < len(nums) && nums[n/2+i] < k; i++ {
			res += k - nums[n/2+i]
		}
	}

	return int64(res)
}
