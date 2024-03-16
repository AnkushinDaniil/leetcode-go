package trappingRainWater

func trap2(height []int) int {
	maxHeight := 0
	bigSum := 0
	smallSum := 0

	for i, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
		bigSum += maxHeight
		smallSum += height[i]
	}

	secondMax := 0

	for i := len(height) - 1; secondMax < maxHeight; i-- {
		if height[i] > secondMax {
			secondMax = height[i]
		}
		bigSum -= maxHeight - secondMax
	}

	return bigSum - smallSum
}

func trap(height []int) int {
	sum := 0
	locMax := 0
	l, r := 0, len(height)-1

	for l < r {
		for height[l] <= locMax && l < r {
			sum += locMax - height[l]
			l++
		}
		for height[r] <= locMax && l < r {
			sum += locMax - height[r]
			r--
		}
		if height[l] < height[r] {
			locMax = height[l]
		} else {
			locMax = height[r]
		}
	}

	return sum
}
