package containsDuplicate

func containsDuplicate(nums []int) bool {
	table := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := table[num]; ok {
			return true
		} else {
			table[num] = struct{}{}
		}
	}

	return false
}
