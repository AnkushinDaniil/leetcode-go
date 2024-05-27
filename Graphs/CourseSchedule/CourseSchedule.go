package courseSchedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	table := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	res := true

	var dfs func(int)
	dfs = func(i int) {
		for len(table[i]) > 0 {
			if visited[table[i][0]] {
				res = false
				return
			}
			visited[table[i][0]] = true
			dfs(table[i][0])
			visited[table[i][0]] = false
			table[i] = table[i][1:]
		}
	}

	for i := 0; i < len(prerequisites); i++ {
		table[prerequisites[i][0]] = append(table[prerequisites[i][0]], prerequisites[i][1])
	}

	for i := 0; i < numCourses && res; i++ {
		dfs(i)
	}

	return res
}
