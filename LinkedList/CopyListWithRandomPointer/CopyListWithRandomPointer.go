package copyListWithRandomPointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	table := make(map[*Node]*Node, 0)
	cur := head

	for cur != nil {
		table[cur] = &Node{}
		cur = cur.Next
	}

	cur = head

	for cur != nil {
		table[cur].Val = cur.Val
		table[cur].Next = table[cur.Next]
		table[cur].Random = table[cur.Random]
		cur = cur.Next
	}

	return table[head]
}
