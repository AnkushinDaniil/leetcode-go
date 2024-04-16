package maximumDepthOfBinaryTree

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
			root:   CreateTreeFromArray([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			output: 3,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{1, nil, 2}),
			output: 2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, maxDepth(testCase.root), testCase.output)
		})
	}
}
