package reverseNodesInKGroup

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		head   *linkedList.ListNode
		k      int
		output *linkedList.ListNode
	}{
		{
			name:   "Example 1",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			k:      2,
			output: linkedList.CreateListFromArray([]int{2, 1, 4, 3, 5}),
		},
		{
			name:   "Example 2",
			head:   linkedList.CreateListFromArray([]int{1, 2, 3, 4, 5}),
			k:      3,
			output: linkedList.CreateListFromArray([]int{3, 2, 1, 4, 5}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, reverseKGroup(testCase.head, testCase.k), testCase.output)
		})
	}
}
