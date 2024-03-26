package binarySearch

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		i := (l + r) / 2
		if nums[i] == target {
			return i
		} else if nums[i] < target {
			l = i + 1
		} else {
			r = i - 1
		}
	}
	return -1
}
