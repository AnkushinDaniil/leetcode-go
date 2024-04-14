package reverseNodesInKGroup

import "github.com/AnkushinDaniil/leetcode-go/linkedList"

type ListNode = linkedList.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	groupPrev := dummy

	for {
		kth := getKth(groupPrev, k)

		if kth == nil {
			break
		}

		groupNext := kth.Next
		prev := kth.Next
		cur := groupPrev.Next

		for cur != groupNext {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}

		groupPrev.Next, groupPrev = kth, groupPrev.Next
	}

	return dummy.Next
}

func getKth(l *ListNode, k int) *ListNode {
	for l != nil && k > 0 {
		l = l.Next
		k--
	}
	return l
}
