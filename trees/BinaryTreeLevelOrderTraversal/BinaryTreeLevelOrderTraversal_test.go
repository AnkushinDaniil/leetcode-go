package binaryTreeLevelOrderTraversal

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
		output [][]int
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			output: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{1}),
			output: [][]int{{1}},
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{}),
			output: [][]int{},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				levelOrder(testCase.root),
				testCase.output,
			)
		})
	}
}
