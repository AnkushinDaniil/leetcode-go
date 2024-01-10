package MaxConsecutiveOnes

func FindMaxConsecutiveOnes(nums []int) int {
	res := 0
	maxRes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			res++
		} else {
			if res > maxRes {
				maxRes = res
			}
			res = 0
		}
	}
	if res > maxRes {
		maxRes = res
	}
	return maxRes
}
