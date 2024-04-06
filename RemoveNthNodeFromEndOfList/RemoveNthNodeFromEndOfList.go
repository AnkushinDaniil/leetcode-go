package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur, prev := head, head

	for ; n > 0; n-- {
		cur = cur.Next
	}

	if cur == nil {
		return head.Next
	}

	for cur.Next != nil {
		cur, prev = cur.Next, prev.Next
	}

	prev.Next = prev.Next.Next

	return head
}
