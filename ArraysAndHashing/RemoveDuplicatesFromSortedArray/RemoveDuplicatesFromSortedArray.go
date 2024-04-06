package removeDuplicatesFromSortedArray

func RemoveDuplicates(nums []int) int {
	k := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			nums[k] = nums[i+1]
			k++
		}
	}
	return k
}
