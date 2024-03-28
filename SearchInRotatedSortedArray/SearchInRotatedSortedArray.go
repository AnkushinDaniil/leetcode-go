package searchInRotatedSortedArray

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if target >= nums[0] {
		for l <= r {
			m := (l + r) / 2
			if nums[m] == target {
				return m
			}
			if nums[m] >= nums[0] {
				if target > nums[m] {
					l = m + 1
				} else {
					r = m - 1
				}
			} else {
				r = m - 1
			}
		}
	} else {
		for l <= r {
			m := (l + r) / 2
			if nums[m] == target {
				return m
			}
			if !(nums[m] >= nums[0]) {
				if target > nums[m] {
					l = m + 1
				} else {
					r = m - 1
				}
			} else {
				l = m + 1
			}
		}
	}

	return -1
}

func search_(nums []int, target int) int {
	l, r := 0, len(nums)-1
	isTargetOnTheLeft := target >= nums[0]

	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if isTargetOnTheLeft == (nums[m] >= nums[0]) {
			if target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		} else if isTargetOnTheLeft {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
