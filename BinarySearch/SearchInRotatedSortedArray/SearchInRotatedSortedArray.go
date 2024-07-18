package searchInRotatedSortedArray

func search(nums []int, target int) int {
	l, r, m := 0, len(nums)-1, 0
	var compare func(int, int) bool
	var update func()

	if target >= nums[0] {
		compare = func(a, b int) bool { return a >= b }
		update = func() { r = m - 1 }
	} else {
		compare = func(a, b int) bool { return a < b }
		update = func() { l = m + 1 }
	}

	for l <= r {
		m = (l + r) / 2
		if nums[m] == target {
			return m
		}
		if compare(nums[m], nums[0]) {
			if target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else {
			update()
		}
	}

	return -1
}
