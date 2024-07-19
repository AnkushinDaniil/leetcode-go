package lowestCommonAncestorOfABinarySearchTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name               string
		root, p, q, output *TreeNode
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:      CreateTreeFromArray([]interface{}{2}),
			q:      CreateTreeFromArray([]interface{}{8}),
			output: CreateTreeFromArray([]interface{}{6}),
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}),
			p:      CreateTreeFromArray([]interface{}{2}),
			q:      CreateTreeFromArray([]interface{}{4}),
			output: CreateTreeFromArray([]interface{}{2}),
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{2, 1}),
			p:      CreateTreeFromArray([]interface{}{2}),
			q:      CreateTreeFromArray([]interface{}{1}),
			output: CreateTreeFromArray([]interface{}{2}),
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}),
			p:      CreateTreeFromArray([]interface{}{5}),
			q:      CreateTreeFromArray([]interface{}{4}),
			output: CreateTreeFromArray([]interface{}{2}),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				lowestCommonAncestor(testCase.root, testCase.p, testCase.q).Val,
				testCase.output.Val,
			)
		})
	}
}
