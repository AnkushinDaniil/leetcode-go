package containerWithMostWater

func maxArea(height []int) int {
	l, r, area := 0, len(height)-1, 0

	for l < r {
		if height[l] < height[r] {
			area = max(area, height[l]*(r-l))
			l++
		} else {
			area = max(area, height[r]*(r-l))
			r--
		}
	}

	return area
}
