package copyListWithRandomPointer

import (
	"github.com/AnkushinDaniil/leetcode-go/linkedList"
)

type ListNode = linkedList.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for l1.Next != nil && l2.Next != nil {
		sum := l1.Val + l2.Val + cur.Val
		cur.Val = sum % 10
		cur.Next = new(ListNode)
		cur.Next.Val += sum / 10
		l1, l2, cur = l1.Next, l2.Next, cur.Next
	}

	sum := l1.Val + l2.Val + cur.Val
	cur.Val = sum % 10
	l1, l2 = l1.Next, l2.Next

	var l *ListNode

	if l1 == nil && l2 == nil {
		if sum >= 10 {
			cur.Next = new(ListNode)
			cur.Next.Val = sum / 10
		}
		return res
	} else if l1 != nil {
		l = l1
	} else {
		l = l2
	}
	cur.Next = new(ListNode)
	cur = cur.Next
	cur.Val = sum / 10

	for l.Next != nil {
		sum := l.Val + cur.Val
		cur.Val = sum % 10
		cur.Next = new(ListNode)
		cur.Next.Val += sum / 10
		l, cur = l.Next, cur.Next
	}

	sum = l.Val + cur.Val
	cur.Val = sum % 10
	if sum >= 10 {
		cur.Next = new(ListNode)
		cur.Next.Val = sum / 10
	}

	return res
}
