package reorderList

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *linkedList.ListNode
		n      int
		output *linkedList.ListNode
	}{
		{
			name:   "Example 1",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			n:      2,
			output: linkedList.CreateListFromArray([]int{1, 2, 3, 5}),
		},
		{
			name:   "Example 2",
			head:   linkedList.CreateListFromArray([]int{1}),
			n:      1,
			output: linkedList.CreateListFromArray([]int{}),
		},
		{
			name:   "Example 3",
			head:   linkedList.CreateListFromArray([]int{1, 2}),
			n:      1,
			output: linkedList.CreateListFromArray([]int{1}),
		},
		{
			name:   "Example 4",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			n:      5,
			output: linkedList.CreateListFromArray([]int{2, 3, 4, 5}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, removeNthFromEnd(testCase.head, testCase.n), testCase.output)
		})
	}
}
