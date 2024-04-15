package invertBinaryTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]int) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		root   *TreeNode
		output *TreeNode
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]int{4, 2, 7, 1, 3, 6, 9}),
			output: CreateTreeFromArray([]int{4, 7, 2, 9, 6, 3, 1}),
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]int{2, 1, 3}),
			output: CreateTreeFromArray([]int{2, 3, 1}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, invertTree(testCase.root), testCase.output)
		})
	}
}
