package jumpGameII

func jump(nums []int) int {
	var steps, maxIndex, index int
	for i := 0; i < len(nums)-1; i++ {
		maxIndex = max(maxIndex, i+nums[i])
		if i == index {
			steps++
			index = maxIndex
		}
	}
	return steps
}

func jumpLCT(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	m := len(nums) - 1
	var steps, canReach, needChoose int
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		if i+val > canReach {
			canReach = i + val
			if canReach >= m {
				return steps + 1
			}
		}
		if i == needChoose {
			needChoose = canReach
			steps++
		}
	}

	return steps
}

func jumpLCM(nums []int) int {
	length := len(nums) - 1
	var maxReachable, lastJumpedPos, jumps int
	for i := 0; i < length; i++ {
		maxReachable = max(maxReachable, i+nums[i])
		if i == lastJumpedPos {
			lastJumpedPos = maxReachable
			jumps++
		}
	}
	return jumps
}
