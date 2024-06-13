package mergeIntervals

import "slices"

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	l, r := intervals[0][0], intervals[0][1]
	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > r {
			intervals[i][0] = l
			intervals[i][1] = r
			i++
			l, r = intervals[j][0], intervals[j][1]
		} else if intervals[j][0] == r || intervals[j][1] > r {
			r = intervals[j][1]
		}
	}
	intervals[i][0] = l
	intervals[i][1] = r

	return intervals[:i+1]
}

func mergeJ(intervals [][]int) [][]int {
	j := 1
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := 0; i+j < len(intervals); i++ {
		intervals[i] = intervals[i+j-1]
		for i+j < len(intervals) && intervals[i][1] >= intervals[i+j][0] {
			if intervals[i+j][1] > intervals[i][1] {
				intervals[i][1] = intervals[i+j][1]
			}
			j++
		}
	}
	if intervals[len(intervals)-j][1] < intervals[len(intervals)-1][0] {
		intervals[len(intervals)-j] = intervals[len(intervals)-1]
	}
	return intervals[:len(intervals)-j+1]
}

func mergeInsert(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals))

	for i := range intervals {
		res = insert(res, intervals[i])
	}

	return res
}

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
