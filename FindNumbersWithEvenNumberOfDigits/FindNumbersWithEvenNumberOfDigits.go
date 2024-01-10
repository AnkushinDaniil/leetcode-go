package FindNumbersWithEvenNumberOfDigits

import "strconv"

func FindNumbers(nums []int) int {
	res := 0
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			res++
		}
	}
	return res
}
