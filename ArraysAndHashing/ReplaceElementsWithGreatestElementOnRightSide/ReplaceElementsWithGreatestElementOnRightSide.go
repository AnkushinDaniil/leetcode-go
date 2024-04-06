package ReplaceElementsWithGreatestElementOnRightSide

func ReplaceElements(arr []int) []int {
	n := len(arr)
	maximum := -1
	for i := n - 1; i >= 0; i-- {
		arr[i], maximum = maximum, max(maximum, arr[i])
	}
	return arr
}

// [17,18,5,4,6,1]
// [18,6,6,6,1,-1]
