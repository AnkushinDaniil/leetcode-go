package redundantConnection

func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges))
	rank := make([]int, len(edges))

	for i := 0; i < len(edges); i++ {
		parents[i] = i
		rank[i] = 0
	}

	for i := 0; i < len(edges); i++ {
		parent := parents[edges[i][0]-1]
		for parent != parents[parent] {
			parent = parents[parent]
		}
		child := parents[edges[i][1]-1]
		for child != parents[child] {
			child = parents[child]
		}
		if parents[parent] == parents[child] {
			return edges[i]
		}
		if rank[parent] >= rank[child] {
			parents[child] = parent
			rank[parent]++
		} else {
			parents[parent] = child
			rank[child]++
		}
	}

	return edges[len(edges)-1]
}
