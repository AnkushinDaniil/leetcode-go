package partitionEqualSubsetSum

func canPartition(nums []int) bool {
	sum := 0
	memo := [10000]bool{}
	memo[0] = true

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1

	for i := 0; i < len(nums); i++ {
		if memo[sum] {
			return true
		} else {
			for j := sum; j >= nums[i]; j-- {
				memo[j] = memo[j] || memo[j-nums[i]]
			}
		}
	}

	return false
}

func canPartitionMap(nums []int) bool {
	sum := 0
	memo := make(map[int]struct{}, len(nums))
	keys := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	if sum&1 == 1 {
		return false
	}

	sum >>= 1

	memo[0] = struct{}{}
	for i := 0; i < len(nums); i++ {
		for key := range memo {
			t := key + nums[i]
			if t < sum {
				if _, ok := memo[t]; !ok {
					keys = append(keys, t)
				}
			} else if t == sum {
				return true
			}
		}
		for j := 0; j < len(keys); j++ {
			memo[keys[j]] = struct{}{}
		}
		keys = keys[0:0]
	}

	return false
}
