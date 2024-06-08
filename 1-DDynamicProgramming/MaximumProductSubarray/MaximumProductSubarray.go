package maximumProductSubarray

func maxProduct(nums []int) int {
	res, maximum, minimum := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		minimum = min(nums[i], nums[i]*minimum)
		maximum = max(nums[i], nums[i]*maximum)
		res = max(maximum, res)
	}
	return res
}
