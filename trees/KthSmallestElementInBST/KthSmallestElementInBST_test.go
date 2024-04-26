package kthSmallestElementInBST

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
		k      int
		output int
	}{
		{
			name:   "Example 1",
			root:   CreateTreeFromArray([]interface{}{3, 1, 4, nil, 2}),
			k:      1,
			output: 1,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{5, 3, 6, 2, 4, nil, nil, 1}),
			k:      3,
			output: 3,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				kthSmallest(testCase.root, testCase.k),
			)
		})
	}
}
