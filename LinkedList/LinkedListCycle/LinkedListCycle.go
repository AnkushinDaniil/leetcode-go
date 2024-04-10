package linkedListCycle

import (
	"github.com/AnkushinDaniil/leetcode-go/linkedList"
)

type ListNode = linkedList.ListNode

func hasCycleSlowFast(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	cur := head
	for cur != nil {
		if cur.Next == head {
			return true
		} else {
			cur, cur.Next = cur.Next, head
		}
	}
	return false
}

func hasCycleTable(head *ListNode) bool {
	table := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := table[head.Next]; ok {
			return true
		} else {
			table[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}
