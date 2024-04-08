package linkedList

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListFromArray(arr []int) *ListNode {
	res := make([]ListNode, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		res[i] = ListNode{
			Val:  arr[i],
			Next: &res[i+1],
		}
	}
	if len(arr) > 0 {
		res[len(arr)-1] = ListNode{
			Val:  arr[len(arr)-1],
			Next: nil,
		}
		return &res[0]
	}

	return nil
}

func (l *ListNode) String() string {
	cur := l
	var res strings.Builder
	for cur != nil {
		res.WriteString(strconv.Itoa(cur.Val) + " ")
		cur = cur.Next
	}
	return res.String()
}
