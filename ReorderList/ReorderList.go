package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow, slow.Next = slow.Next, nil

	var prev *ListNode = nil

	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	cur := head

	for cur.Next != nil {
		next := cur.Next
		cur.Next = prev
		prev = prev.Next
		cur.Next.Next = next
		cur = cur.Next.Next
	}

	cur.Next = prev
}
