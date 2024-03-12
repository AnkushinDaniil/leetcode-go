package longestConsecutiveSequence

func longestConsecutive(nums []int) int {
	table := make(map[int]struct{}, len(nums))
	res := 0

	for _, num := range nums {
		table[num] = struct{}{}
	}

	var length int

	for num := range table {
		if _, ok := table[num-1]; !ok {
			length = 1

			for _, ok := table[num+length]; ok; _, ok = table[num+length] {
				length++
			}

			res = max(res, length)
		}
	}

	return res
}
