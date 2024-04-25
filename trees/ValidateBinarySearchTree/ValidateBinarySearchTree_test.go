package validateBinarySearchTree

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
			root:   CreateTreeFromArray([]interface{}{2, 1, 3}),
			output: true,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{5, 1, 4, nil, nil, 3, 6}),
			output: false,
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{5, 4, 6, nil, nil, 3, 7}),
			output: false,
		},
		{
			name:   "Example 4",
			root:   CreateTreeFromArray([]interface{}{2, 2, 2}),
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				isValidBST(testCase.root),
			)
		})
	}
}
