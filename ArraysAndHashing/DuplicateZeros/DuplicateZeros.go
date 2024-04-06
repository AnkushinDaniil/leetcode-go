package DuplicateZeros

func DuplicateZeros(nums []int) []int {
	n := len(nums)
	zeroCounter := 0
	nonZeroCounter := 0
	for i := 0; nonZeroCounter+2*zeroCounter < n; i++ {
		if nums[i] == 0 {
			zeroCounter++
		} else {
			nonZeroCounter++
		}
	}
	if nonZeroCounter+2*zeroCounter > n {
		nums[n-1] = 0
		n--
		zeroCounter--
	}
	for i := n - 1; i-zeroCounter >= 0; i-- {
		if nums[i-zeroCounter] == 0 {
			nums[i] = 0
			nums[i-1] = 0
			zeroCounter--
			i--
		} else {
			nums[i] = nums[i-zeroCounter]
		}
	}
	return nums
}
