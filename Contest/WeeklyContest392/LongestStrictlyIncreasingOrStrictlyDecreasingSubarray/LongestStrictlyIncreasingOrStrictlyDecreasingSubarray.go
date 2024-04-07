package binarySearch

func longestMonotonicSubarray(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	l, r, result := 0, 1, 0
	isIncrease := nums[1] > nums[0]

	for ; r < len(nums); r++ {
		if nums[r] == nums[r-1] {
			if t := r - l; t > result {
				result = t
			}
			l = r
			if r < len(nums)-1 {
				isIncrease = nums[r] < nums[r+1]
			}
		} else if isIncrease != (nums[r] > nums[r-1]) {
			if t := r - l; t > result {
				result = t
			}
			l = r - 1
			isIncrease = !isIncrease
		}
	}

	if t := r - l; t > result {
		result = t
	}

	return result
}
