package reorderList

func findDuplicate(nums []int) int {
	long := make([]byte, len(nums))
	for _, num := range nums {
		firstByte, secondByte := num>>3, num&7
		var t byte = 1 << secondByte
		if long[firstByte]&t > 0 {
			return num
		} else {
			long[firstByte] |= t
		}
	}
	return -1
}

func findDuplicate_(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = nums[0]
	slow = nums[slow]

	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return fast
}
