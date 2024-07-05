package missingNumber

func missingNumber(nums []int) int {
	sum := (len(nums) * (len(nums) + 1)) >> 1
	for i := range nums {
		sum -= nums[i]
	}
	return sum
}
