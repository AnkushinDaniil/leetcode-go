package threeSum

func threeSum(nums []int) [][]int {
	quickSortAlgorithm(&nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, l, r := -nums[i], i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] > target {
				r--
			} else if nums[l]+nums[r] < target {
				l++
			} else {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
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
