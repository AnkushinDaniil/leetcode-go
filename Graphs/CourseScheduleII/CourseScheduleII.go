package courseScheduleII

func findOrder(numCourses int, prerequisites [][]int) []int {
	table := make([][]int, numCourses)
	visited := make([]bool, numCourses)
	res := make([]int, 0, numCourses)
	topology := true

	var dfs func(int)
	dfs = func(i int) {
		if table[i] == nil {
			return
		}
		for len(table[i]) > 0 {
			if visited[table[i][0]] {
				topology = false
				return
			}
			visited[table[i][0]] = true
			dfs(table[i][0])
			visited[table[i][0]] = false
			table[i] = table[i][1:]
		}
		res = append(res, i)
		table[i] = nil
	}

	for i := 0; i < numCourses && topology; i++ {
		table[i] = []int{}
	}

	for i := 0; i < len(prerequisites); i++ {
		table[prerequisites[i][0]] = append(table[prerequisites[i][0]], prerequisites[i][1])
	}

	for i := 0; i < numCourses && topology; i++ {
		dfs(i)
	}

	if topology {
		return res
	} else {
		return []int{}
	}
}
