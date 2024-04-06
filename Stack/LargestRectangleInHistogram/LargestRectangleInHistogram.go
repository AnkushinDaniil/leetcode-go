package largestRectangleInHistogram

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 1, n)
	maxArea := heights[0]
	stack[0] = 0

	for i := 1; i < n; i++ {
		if heights[i] >= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for j := 1; len(stack) >= j && heights[i] < heights[stack[len(stack)-j]]; j++ {
				maxArea = max(maxArea, (i-stack[len(stack)-j])*heights[stack[len(stack)-j]])
				heights[stack[len(stack)-j]] = heights[i]
			}
		}
	}

	for len(stack) > 0 {
		maxArea = max(maxArea, (n-stack[len(stack)-1])*heights[stack[len(stack)-1]])
		stack = stack[:len(stack)-1]
	}

	return maxArea
}
