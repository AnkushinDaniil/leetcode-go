package mergeTwoSortedLists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		list1  *ListNode
		list2  *ListNode
		output *ListNode
	}{
		{
			name:   "Example 1",
			list1:  linkedListFromArray([]int{1, 2, 4}),
			list2:  linkedListFromArray([]int{1, 3, 4}),
			output: linkedListFromArray([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:   "Example 2",
			list1:  linkedListFromArray([]int{}),
			list2:  linkedListFromArray([]int{}),
			output: linkedListFromArray([]int{}),
		},
		{
			name:   "Example 3",
			list1:  linkedListFromArray([]int{}),
			list2:  linkedListFromArray([]int{0}),
			output: linkedListFromArray([]int{0}),
		},
		{
			name:   "Example 4",
			list1:  linkedListFromArray([]int{1}),
			list2:  linkedListFromArray([]int{}),
			output: linkedListFromArray([]int{1}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, mergeTwoLists(testCase.list1, testCase.list2), testCase.output)
		})
	}
}

func linkedListFromArray(arr []int) *ListNode {
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
