package countGoodNodesInBinaryTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		root   *TreeNode
		output int
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{3, 1, 4, 3, nil, 1, 5}),
			output: 4,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{3, 3, nil, 4, 2}),
			output: 3,
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{}),
			output: 0,
		},
		{
			name:   "Example 4",
			root:   CreateTreeFromArray([]interface{}{1}),
			output: 1,
		},
		{
			name:   "Example 5",
			root:   CreateTreeFromArray([]interface{}{2, nil, 4, 10, 8, nil, nil, 4}),
			output: 4,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				goodNodes(testCase.root),
			)
		})
	}
}
