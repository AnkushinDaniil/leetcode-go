package sortArrayByParity

func sortArrayByParity(nums []int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		if nums[i]&1 == 0 {
			i++
		} else if nums[j]&1 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			j--
		}
	}

	return nums
}
