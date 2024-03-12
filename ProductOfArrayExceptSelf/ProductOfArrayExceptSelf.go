package productOfArrayExceptSelf

func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	res[0] = nums[0]

	for i := 1; i < n-1; i++ {
		res[i] = nums[i] * res[i-1]
	}

	product := 1

	for i := n - 1; i > 0; i-- {
		res[i] = product * res[i-1]
		product *= nums[i]
	}

	res[0] = product

	return res
}
