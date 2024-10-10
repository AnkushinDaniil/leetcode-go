package sortanarray

func quickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	i := -1
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] < nums[len(nums)-1] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i]

	quickSort(nums[:i])
	quickSort(nums[i+1:])
}
