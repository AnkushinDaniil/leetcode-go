package ValidMountainArray

func ValidMountainArray(arr []int) bool {
	i := 0
	for ; i+1 < len(arr) && arr[i] < arr[i+1]; i++ {
	}
	if i == 0 || i == len(arr)-1 {
		return false
	}
	for ; i+1 < len(arr) && arr[i] > arr[i+1]; i++ {
	}
	return i == len(arr)-1
}
