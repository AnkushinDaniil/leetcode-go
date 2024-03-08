package heightChecker

func heightChecker(heights []int) int {
	res := 0
	sortedHeights := make([]int, len(heights))
	copy(sortedHeights, heights)

	quickSortAlgorithm(&sortedHeights)

	for i := 0; i < len(heights); i++ {
		if heights[i] != sortedHeights[i] {
			res++
		}
	}

	return res
}

func quickSortAlgorithm(arr *[]int) {
	quickSort(arr, 0, len(*arr)-1)
}

func quickSort(arr *[]int, l, r int) {
	var i int
	if l < r {
		i = partition(arr, l, r)
		quickSort(arr, l, i-1)
		quickSort(arr, i, r)
	}
}

func partition(arr *[]int, l, r int) int {
	p := (*arr)[(l+r)/2]
	for l <= r {
		for (*arr)[l] < p {
			l++
		}
		for (*arr)[r] > p {
			r--
		}
		if l <= r {
			(*arr)[r], (*arr)[l] = (*arr)[l], (*arr)[r]
			l++
			r--
		}
	}
	return l
}
