package floodFill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc] + 0
	if originalColor == color {
		return image
	}
	stack := make([][2]int, 0, len(image)*len(image[0]))
	stack = append(stack, [2]int{sr, sc})

	for len(stack) > 0 {
		i := stack[len(stack)-1][0]
		j := stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		if image[i][j] == originalColor {
			image[i][j] = color
			if i-1 >= 0 && image[i-1][j] != color {
				stack = append(stack, [2]int{i - 1, j})
			}
			if j-1 >= 0 && image[i][j-1] != color {
				stack = append(stack, [2]int{i, j - 1})
			}
			if i+1 < len(image) && image[i+1][j] != color {
				stack = append(stack, [2]int{i + 1, j})
			}
			if j+1 < len(image[0]) && image[i][j+1] != color {
				stack = append(stack, [2]int{i, j + 1})
			}
		}
	}

	return image
}

func floodFillQueue(image [][]int, sr int, sc int, color int) [][]int {
	originalColor := image[sr][sc] + 0
	if originalColor == color {
		return image
	}
	queue := make([][2]int, 0, len(image)*len(image[0]))
	queue = append(queue, [2]int{sr, sc})

	for len(queue) > 0 {
		i := queue[0][0]
		j := queue[0][1]
		queue = queue[1:]
		if image[i][j] == originalColor {
			image[i][j] = color
			if i-1 >= 0 && image[i-1][j] != color {
				queue = append(queue, [2]int{i - 1, j})
			}
			if j-1 >= 0 && image[i][j-1] != color {
				queue = append(queue, [2]int{i, j - 1})
			}
			if i+1 < len(image) && image[i+1][j] != color {
				queue = append(queue, [2]int{i + 1, j})
			}
			if j+1 < len(image[0]) && image[i][j+1] != color {
				queue = append(queue, [2]int{i, j + 1})
			}
		}
	}

	return image
}
