package jumpGame

func canJump(nums []int) bool {
	maxIndex := nums[0]
	for i := 0; i < len(nums); i++ {
		if i > maxIndex {
			return false
		}
		if maxIndex < i+nums[i] {
			maxIndex = i + nums[i]
		}
	}
	return true
}
