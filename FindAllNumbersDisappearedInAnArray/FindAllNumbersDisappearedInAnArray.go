package findAllNumbersDisappearedInAnArray

func findDisappearedNumbers_hash(nums []int) []int {
	n := len(nums)
	table := make(map[int]bool, n)

	for _, num := range nums {
		table[num] = true
	}

	nums = make([]int, 0, n)

	for i := 1; i <= n; i++ {
		if ok, _ := table[i]; !ok {
			nums = append(nums, i)
		}
	}

	return nums
}

func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	k := 0

	for i, num := range nums {
		if num-1 != i {
			nums[k] = i + 1
			k++
		}
	}

	return nums[:k]
}
