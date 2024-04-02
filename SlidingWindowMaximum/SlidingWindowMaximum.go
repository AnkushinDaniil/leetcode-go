package slidingWindowMaximum

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	res := make([]int, n-k+1)
	deque := make([]int, 0)
	i := 0

	calc := func() {
		if len(deque) > 0 && i-k == deque[0] {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
	}

	for ; i < k-1; i++ {
		calc()
	}

	for ; i < n; i++ {
		calc()

		res[i-k+1] = nums[deque[0]]
	}

	return res
}
