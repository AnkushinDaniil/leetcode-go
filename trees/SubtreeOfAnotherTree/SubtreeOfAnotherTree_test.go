package subtreeOfAnotherTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name    string
		root    *TreeNode
		subRoot *TreeNode
		output  bool
	}{
		{
			name:    "Example 1",
			root:    CreateTreeFromArray([]interface{}{3, 4, 5, 1, 2}),
			subRoot: CreateTreeFromArray([]interface{}{4, 1, 2}),
			output:  true,
		},
		{
			name:    "Example 2",
			root:    CreateTreeFromArray([]interface{}{3, 4, 5, 1, 2, nil, nil, nil, nil, 0}),
			subRoot: CreateTreeFromArray([]interface{}{4, 1, 2}),
			output:  false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isSubtree(testCase.root, testCase.subRoot), testCase.output)
		})
	}
}
