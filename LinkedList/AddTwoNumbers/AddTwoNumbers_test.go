package copyListWithRandomPointer

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		l1     *ListNode
		l2     *ListNode
		output *ListNode
	}{
		{
			name:   "Example 1",
			l1:     linkedList.CreateListFromArray([]int{2, 4, 3}),
			l2:     linkedList.CreateListFromArray([]int{5, 6, 4}),
			output: linkedList.CreateListFromArray([]int{7, 0, 8}),
		},
		{
			name:   "Example 2",
			l1:     linkedList.CreateListFromArray([]int{0}),
			l2:     linkedList.CreateListFromArray([]int{0}),
			output: linkedList.CreateListFromArray([]int{0}),
		},
		{
			name:   "Example 3",
			l1:     linkedList.CreateListFromArray([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:     linkedList.CreateListFromArray([]int{9, 9, 9, 9}),
			output: linkedList.CreateListFromArray([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name:   "Example 4",
			l1:     linkedList.CreateListFromArray([]int{9, 9, 9, 9}),
			l2:     linkedList.CreateListFromArray([]int{9, 9, 9, 9}),
			output: linkedList.CreateListFromArray([]int{8, 9, 9, 9, 1}),
		},
		{
			name:   "Example 5",
			l1:     linkedList.CreateListFromArray([]int{9, 9, 1}),
			l2:     linkedList.CreateListFromArray([]int{1}),
			output: linkedList.CreateListFromArray([]int{0, 0, 2}),
		},
		{
			name:   "Example 6",
			l1:     linkedList.CreateListFromArray([]int{9, 1, 6}),
			l2:     linkedList.CreateListFromArray([]int{0}),
			output: linkedList.CreateListFromArray([]int{9, 1, 6}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, addTwoNumbers(testCase.l1, testCase.l2), testCase.output)
		})
	}
}
