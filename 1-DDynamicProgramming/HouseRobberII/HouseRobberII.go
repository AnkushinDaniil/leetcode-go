package houseRobberII

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	rob1, rob2 := 0, 0
	rob3, rob4 := 0, 0

	for i := 0; i < len(nums)-1; i++ {
		rob1, rob2 = rob2, max(rob1+nums[i], rob2)
		rob3, rob4 = rob4, max(rob3+nums[i+1], rob4)
	}

	return max(rob2, rob4)
}
