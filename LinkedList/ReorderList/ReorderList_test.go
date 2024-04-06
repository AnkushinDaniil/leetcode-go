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
		output *linkedList.ListNode
	}{
		{
			name:   "Example 1",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4}),
			output: linkedList.CreateListFromArray([]int{1, 4, 2, 3}),
		},
		{
			name:   "Example 2",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			output: linkedList.CreateListFromArray([]int{1, 5, 2, 4, 3}),
		},
		{
			name:   "Example 3",
			head:   linkedList.CreateListFromArray([]int{1}),
			output: linkedList.CreateListFromArray([]int{1}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			reorderList(testCase.head)
			assert.Equal(t, testCase.head, testCase.output)
		})
	}
}
