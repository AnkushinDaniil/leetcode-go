package reverseLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		t := head.Next
		head.Next = prev
		prev = head
		head = t
	}
	return prev
}
