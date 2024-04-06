package reverseLinkedList

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *linkedList.ListNode
		output *linkedList.ListNode
	}{
		{
			name:   "Example 1",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			output: linkedList.CreateListFromArray([]int{5, 4, 3, 2, 1}),
		},
		{
			name:   "Example 2",
			head:   linkedList.CreateListFromArray([]int{1, 2}),
			output: linkedList.CreateListFromArray([]int{2, 1}),
		},
		{
			name:   "Example 3",
			head:   linkedList.CreateListFromArray([]int{}),
			output: linkedList.CreateListFromArray([]int{}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, reverseList(testCase.head), testCase.output)
		})
	}
}
