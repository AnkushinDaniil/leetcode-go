package partitionLabels

import "slices"

func partitionLabels(s string) []int {
	arrEnd := [26]int{}
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		arrEnd[s[i]-'a'] = i
	}

	res := make([]int, 0, 8)

	for i := 0; i < len(s); i++ {
		if arrEnd[s[i]-'a'] > end {
			end = arrEnd[s[i]-'a']
		}
		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func partitionLabelsHashMap(s string) []int {
	hashMap := make(map[byte]int)
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		hashMap[s[i]] = i
	}

	res := make([]int, 0, 4)

	for i := 0; i < len(s); i++ {
		if hashMap[s[i]] > end {
			end = hashMap[s[i]]
		}
		if end == i {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

func partitionLabelsIntervals(s string) []int {
	intervalsArr := [26][2]int{}
	for i := 0; i < 26; i++ {
		intervalsArr[i][0] = -1
	}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if intervalsArr[index][0] == -1 {
			intervalsArr[index] = [2]int{i, i}
		} else {
			intervalsArr[index][1] = i
		}
	}
	counter := 0
	for i := 0; i < 26; i++ {
		if intervalsArr[i][0] == -1 {
			continue
		} else {
			intervalsArr[counter] = intervalsArr[i]
			counter++
		}
	}

	intervals := intervalsArr[:counter]

	slices.SortFunc(intervals, func(a, b [2]int) int {
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

	res := make([]int, i+1)
	for i := 0; i < len(res); i++ {
		res[i] = intervals[i][1] - intervals[i][0] + 1
	}
	return res
}
