package insertInterval

func insert(intervals [][]int, newInterval []int) [][]int {
	i := 0
	for ; i < len(intervals) && newInterval[0] > intervals[i][1]; i++ {
	}
	if i == len(intervals) {
		return append(intervals, newInterval)
	} else if newInterval[1] < intervals[i][0] {
		intervals = append(intervals, []int{})
		for j := len(intervals) - 1; j >= i+1; j-- {
			intervals[j] = intervals[j-1]
		}
		intervals[i] = newInterval
		return intervals
	} else {
		j := i
		for j+1 < len(intervals) && newInterval[1] >= intervals[j+1][0] {
			j++
		}
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[j][1])
		intervals = append(intervals[:i], intervals[j:]...)
		intervals[i] = newInterval
		return intervals
	}
}

func insertOld(intervals [][]int, newInterval []int) [][]int {
	for i := range intervals {
		if newInterval[1] < intervals[i][0] {
			intervals = append(intervals, []int{})
			for j := len(intervals) - 1; j >= i+1; j-- {
				intervals[j] = intervals[j-1]
			}
			intervals[i] = newInterval
			return intervals
		} else if newInterval[0] <= intervals[i][1] {
			j := i
			for j+1 < len(intervals) && newInterval[1] >= intervals[j+1][0] {
				j++
			}
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[j][1])
			intervals = append(intervals[:i], intervals[j:]...)
			intervals[i] = newInterval
			return intervals
		}
	}
	return append(intervals, newInterval)
}

func insertLCT(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}

	start := newInterval[0]
	end := newInterval[1]
	index := 0
	for index < len(intervals) && start > intervals[index][0] {
		result = append(result, intervals[index])
		index++
	}

	if len(result) > 0 && result[len(result)-1][1] >= start {
		for index < len(intervals) && intervals[index][0] <= end {
			index++
		}
		if intervals[index-1][1] < end {
			result[len(result)-1][1] = end
		} else {
			result[len(result)-1][1] = intervals[index-1][1]
		}
	} else {
		for index < len(intervals) && intervals[index][0] <= end {
			index++
		}
		if index > 0 {
			if intervals[index-1][1] < end {
				result = append(result, []int{start, end})
			} else {
				result = append(result, []int{start, intervals[index-1][1]})
			}
		} else {
			result = append(result, []int{start, end})
		}
	}

	for index < len(intervals) {
		result = append(result, intervals[index])
		index++
	}

	return result
}

func insertLCM(intervals [][]int, newInterval []int) [][]int {
	for i := 0; i < len(intervals); i++ {
		v := intervals[i]
		if v[1] < newInterval[0] {
			continue
		}
		if v[0] > newInterval[1] {
			result := make([][]int, 0, len(intervals))
			if i > 0 {
				result = append(result, intervals[:i]...)
			}

			result = append(result, newInterval)
			result = append(result, intervals[i:]...)
			return result
		}

		if v[0] > newInterval[0] {
			v[0] = newInterval[0]
		}
		j := i + 1
		for ; j < len(intervals); j++ {
			if intervals[j][1] > newInterval[1] {
				break
			}
		}
		if j == len(intervals) {
			if newInterval[1] > v[1] {
				v[1] = newInterval[1]
			}
			intervals = intervals[:i+1]
			return intervals
		}
		if intervals[j][0] > newInterval[1] {
			if newInterval[1] > v[1] {
				v[1] = newInterval[1]
			}
			result := intervals[:i+1]
			result = append(result, intervals[j:]...)
			return result
		}
		v[1] = intervals[j][1]
		result := intervals[:i+1]
		if j+1 < len(intervals) {
			result = append(result, intervals[j+1:]...)
		}

		return result
	}
	intervals = append(intervals, newInterval)
	return intervals
}

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	binarySearch := func(target int) int {
// 		l, r := 0, len(intervals)<<1
// 		for l < r {
// 			m1 := (l + r) >> 2
// 			m2 := (l + r) & 1
// 			if intervals[m1][m2] == target {
// 				return m1<<1 + 1
// 			} else if intervals[m1][m2] < target {
// 				l = m1<<1 + m2 + 1
// 			} else {
// 				r = m1<<1 + m2
// 			}
// 		}
// 		return l
// 	}

// 	index1 := binarySearch(newInterval[0])
// 	index2 := binarySearch(newInterval[1])

// 	i1 := index1 >> 1
// 	i2 := index1 & 1
// 	j1 := index2 >> 1
// 	j2 := index2 & 1
// 	if i2 == 0 && j2 == 0 && i1 == j1 {
// 		intervals = append(intervals, []int{})
// 		for i := len(intervals) - 1; i >= i1+1; i-- {
// 			intervals[i] = intervals[i-1]
// 		}
// 		intervals[i1] = newInterval
// 	} else {
// 		if i2 == 1 {
// 			newInterval[0] = min(intervals[i1][0], newInterval[0])
// 		}
// 		if j2 == 1 {
// 			newInterval[1] = max(intervals[j1][1], newInterval[1])
// 		}
// 		intervals = append(intervals[:i1+i2], intervals[j1+j2:]...)
// 		if len(intervals) == i1 {
// 			intervals = append(intervals, newInterval)
// 		} else {
// 			intervals[i1] = newInterval
// 		}
// 	}
// 	return intervals
// }
