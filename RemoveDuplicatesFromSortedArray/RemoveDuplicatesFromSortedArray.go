package removeDuplicatesFromSortedArray

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if nums[0] == nums[n-1] {
		return 1
	}
	i, j, num := 1, 0, nums[0]
	for i+j < n {
		for num == nums[i+j] {
			j++
			if i+j >= n {
				return i
			}
		}
		nums[i] = nums[i+j]
		num = nums[i+j]
		i++
	}
	return i
}
