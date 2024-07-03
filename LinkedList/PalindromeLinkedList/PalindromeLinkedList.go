package palindromeLinkedList

import "github.com/AnkushinDaniil/leetcode-go/linkedList"

type ListNode = linkedList.ListNode

func isPalindrome(head *ListNode) bool {
	l, r := head, head
	var prev, t *ListNode
	for r != nil {
		r = r.Next
		t = l.Next
		l.Next = prev
		prev = l
		l = t
		if r != nil {
			r = r.Next
		} else {
			prev = prev.Next
		}
	}
	for l != nil {
		if l.Val != prev.Val {
			return false
		}
		l = l.Next
		prev = prev.Next
	}
	return true
}
