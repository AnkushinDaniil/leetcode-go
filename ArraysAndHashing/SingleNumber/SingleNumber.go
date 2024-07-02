package singleNumber

func singleNumber(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i+1] ^= nums[i]
	}
	return nums[len(nums)-1]
}
