package checkIfNAndItsDoubleExist

func CheckIfExist(arr []int) bool {
	checked := make(map[int]bool)
	for _, num := range arr {
		if checked[2*num] || num&1 == 0 && checked[num/2] {
			return true
		}
		checked[num] = true
	}
	return false
}
