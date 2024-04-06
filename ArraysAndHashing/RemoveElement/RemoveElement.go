package removeElement

func RemoveElement(nums []int, val int) int {
	counter := 0
	n := len(nums)
	for i := 0; i+counter < n; i++ {
		for ; i+counter < n && nums[i+counter] == val; counter++ {
		}
		if i+counter < n {
			nums[i] = nums[i+counter]
		}
	}
	return n - counter
}
