package balancedBinaryTree

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
		output bool
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			output: true,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4}),
			output: false,
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{}),
			output: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isBalanced(testCase.root), testCase.output)
		})
	}
}
