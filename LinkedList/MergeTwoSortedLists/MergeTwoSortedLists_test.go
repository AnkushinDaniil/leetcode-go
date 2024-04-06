package mergeTwoSortedLists

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
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
			list1:  linkedList.CreateListFromArray([]int{1, 2, 4}),
			list2:  linkedList.CreateListFromArray([]int{1, 3, 4}),
			output: linkedList.CreateListFromArray([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			name:   "Example 2",
			list1:  linkedList.CreateListFromArray([]int{}),
			list2:  linkedList.CreateListFromArray([]int{}),
			output: linkedList.CreateListFromArray([]int{}),
		},
		{
			name:   "Example 3",
			list1:  linkedList.CreateListFromArray([]int{}),
			list2:  linkedList.CreateListFromArray([]int{0}),
			output: linkedList.CreateListFromArray([]int{0}),
		},
		{
			name:   "Example 4",
			list1:  linkedList.CreateListFromArray([]int{1}),
			list2:  linkedList.CreateListFromArray([]int{}),
			output: linkedList.CreateListFromArray([]int{1}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, mergeTwoLists(testCase.list1, testCase.list2), testCase.output)
		})
	}
}
