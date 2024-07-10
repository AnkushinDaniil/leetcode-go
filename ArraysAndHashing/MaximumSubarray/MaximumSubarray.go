package maximumSubarray

func maxSubArray(nums []int) int {
	sum := nums[0]
	maxSum := nums[0]
	for _, num := range nums[1:] {
		sum = max(sum+num, num)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}
