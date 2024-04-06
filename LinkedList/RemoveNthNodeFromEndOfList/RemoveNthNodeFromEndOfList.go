package reorderList

import "github.com/AnkushinDaniil/leetcode-go/linkedList"

func removeNthFromEnd(head *linkedList.ListNode, n int) *linkedList.ListNode {
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
