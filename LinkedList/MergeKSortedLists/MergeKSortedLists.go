package mergeKSortedLists

import "github.com/AnkushinDaniil/leetcode-go/linkedList"

type ListNode = linkedList.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)

	if n == 0 {
		return nil
	}

	for k := 1; k < n; k *= 2 {
		for i := 0; i+k < n; i += 2 * k {
			lists[i] = mergeTwoLists(lists[i], lists[i+k])
		}
	}

	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	for {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	return res.Next
}
