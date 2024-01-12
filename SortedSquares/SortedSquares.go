package SortedSquares

func SortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	for l <= r {
		if Abs(nums[l]) > Abs(nums[r]) {
			res[r-l] = nums[l] * nums[l]
			l++
		} else {
			res[r-l] = nums[r] * nums[r]
			r--
		}
	}
	return res
}

func Abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
