package maximumSubarray

func maxSubArray(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] > sum {
			sum = nums[i]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}
