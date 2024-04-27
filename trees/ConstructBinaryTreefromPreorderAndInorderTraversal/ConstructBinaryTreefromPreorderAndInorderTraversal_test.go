package constructBinaryTreefromPreorderAndInorderTraversal

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name              string
		preorder, inorder []int
		output            *TreeNode
	}{
		{
			name:     "Example 1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			output:   CreateTreeFromArray([]interface{}{3, 9, 20, nil, nil, 15, 7}),
		},
		{
			name:     "Example 2",
			preorder: []int{-1},
			inorder:  []int{-1},
			output:   CreateTreeFromArray([]interface{}{-1}),
		},
		{
			name:     "Example 3",
			preorder: []int{1, 2},
			inorder:  []int{2, 1},
			output:   CreateTreeFromArray([]interface{}{1, 2}),
		},
		{
			name:     "Example 4",
			preorder: []int{1, 2},
			inorder:  []int{1, 2},
			output:   CreateTreeFromArray([]interface{}{1, nil, 2}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				buildTree(testCase.preorder, testCase.inorder),
			)
		})
	}
}
