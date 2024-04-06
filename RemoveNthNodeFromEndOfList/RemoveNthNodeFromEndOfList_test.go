package reorderList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *ListNode
		n      int
		output *ListNode
	}{
		{
			name:   "Example 1",
			head:   linkedListFromArray([]int{1, 2, 3, 4, 5}),
			n:      2,
			output: linkedListFromArray([]int{1, 2, 3, 5}),
		},
		{
			name:   "Example 2",
			head:   linkedListFromArray([]int{1}),
			n:      1,
			output: linkedListFromArray([]int{}),
		},
		{
			name:   "Example 3",
			head:   linkedListFromArray([]int{1, 2}),
			n:      1,
			output: linkedListFromArray([]int{1}),
		},
		{
			name:   "Example 4",
			head:   linkedListFromArray([]int{1, 2, 3, 4, 5}),
			n:      5,
			output: linkedListFromArray([]int{2, 3, 4, 5}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, removeNthFromEnd(testCase.head, testCase.n), testCase.output)
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
