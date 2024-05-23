package сloneGraph

type Node struct {
	Val       int
	Neighbors []*Node
}

const NODE_TABLE_INITIAL_CAPACITY = 101

func cloneGraphDeque(node *Node) *Node {
	if node == nil {
		return nil
	}
	table := make(map[*Node]*Node, NODE_TABLE_INITIAL_CAPACITY)
	deque := make([]*Node, 0, NODE_TABLE_INITIAL_CAPACITY)
	deque = append(deque, node)
	for len(deque) > 0 {
		cur := deque[0]
		deque = deque[1:]
		if _, ok := table[cur]; !ok {
			table[cur] = &Node{
				Val:       cur.Val,
				Neighbors: make([]*Node, len(cur.Neighbors)),
			}
			for i := 0; i < len(cur.Neighbors); i++ {
				deque = append(deque, cur.Neighbors[i])
			}
		}
	}

	for key, val := range table {
		for i := 0; i < len(key.Neighbors); i++ {
			val.Neighbors[i] = table[key.Neighbors[i]]
		}
	}
	return table[node]
}

func cloneGraphQueue(node *Node) *Node {
	if node == nil {
		return nil
	}
	table := make(map[*Node]*Node, NODE_TABLE_INITIAL_CAPACITY)
	queue := make([]*Node, 0, NODE_TABLE_INITIAL_CAPACITY)
	queue = append(queue, node)
	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if _, ok := table[cur]; !ok {
			table[cur] = &Node{
				Val:       cur.Val,
				Neighbors: make([]*Node, len(cur.Neighbors)),
			}
			for i := 0; i < len(cur.Neighbors); i++ {
				queue = append(queue, cur.Neighbors[i])
			}
		}
	}

	for key, val := range table {
		for i := 0; i < len(key.Neighbors); i++ {
			val.Neighbors[i] = table[key.Neighbors[i]]
		}
	}
	return table[node]
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	table := make(map[*Node]*Node, NODE_TABLE_INITIAL_CAPACITY)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		if _, ok := table[n]; !ok {
			table[n] = &Node{
				Val:       n.Val,
				Neighbors: make([]*Node, len(n.Neighbors)),
			}
			for i := 0; i < len(n.Neighbors); i++ {
				table[n].Neighbors[i] = dfs(n.Neighbors[i])
			}
		}
		return table[n]
	}

	return dfs(node)
}
