package mergeIntervals

import "slices"

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	res := 0
	r := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= r {
			r = intervals[i][1]
		} else {
			res++
			if intervals[i][1] < r {
				r = intervals[i][1]
			}
		}
	}
	return res
}
