package binaryTreeRightSideView

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
		output []int
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{1, 2, 3, nil, 5, nil, 4}),
			output: []int{1, 3, 4},
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{1, nil, 3}),
			output: []int{1, 3},
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{}),
			output: []int{},
		},
		{
			name:   "Example 4",
			root:   CreateTreeFromArray([]interface{}{1, 2}),
			output: []int{1, 2},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				rightSideView(testCase.root),
				testCase.output,
			)
		})
	}
}
