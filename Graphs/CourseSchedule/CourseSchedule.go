package courseSchedule

func canFinish(numCourses int, pre [][]int) bool {
	preMap := make([][]int, numCourses)
	visited := make([]bool, numCourses)

	for i := range pre {
		preMap[pre[i][0]] = append(preMap[pre[i][0]], pre[i][1])
	}

	var dfs func(int) bool
	dfs = func(i int) bool {
		for len(preMap[i]) > 0 {
			if visited[preMap[i][0]] == true {
				return false
			}
			visited[preMap[i][0]] = true
			if !dfs(preMap[i][0]) {
				return false
			}
			visited[preMap[i][0]] = false
			preMap[i] = preMap[i][1:]
		}
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
