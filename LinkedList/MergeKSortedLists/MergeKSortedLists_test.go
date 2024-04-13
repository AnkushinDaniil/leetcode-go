package mergeKSortedLists

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		lists  []*linkedList.ListNode
		output *linkedList.ListNode
	}{
		{
			name:   "Example 1",
			lists:  CreateArrayOfLists([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}),
			output: linkedList.CreateListFromArray([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name:   "Example 2",
			lists:  CreateArrayOfLists([][]int{}),
			output: linkedList.CreateListFromArray([]int{}),
		},
		{
			name:   "Example 3",
			lists:  CreateArrayOfLists([][]int{{}}),
			output: linkedList.CreateListFromArray([]int{}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, mergeKLists(testCase.lists), testCase.output)
		})
	}
}

func CreateArrayOfLists(arrays [][]int) []*ListNode {
	res := make([]*ListNode, len(arrays))
	for i, arr := range arrays {
		res[i] = linkedList.CreateListFromArray(arr)
	}
	return res
}
