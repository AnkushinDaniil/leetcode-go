package houseRobber

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if nums[0] > nums[1] {
		nums[1] = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}

	return nums[len(nums)-1]
}
