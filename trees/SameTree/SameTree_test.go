package sameTree

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		p      *TreeNode
		q      *TreeNode
		output bool
	}{
		{
			name:   "Example 1",
			p:      CreateTreeFromArray([]interface{}{1, 2, 3}),
			q:      CreateTreeFromArray([]interface{}{1, 2, 3}),
			output: true,
		},
		{
			name:   "Example 2",
			p:      CreateTreeFromArray([]interface{}{1, 2}),
			q:      CreateTreeFromArray([]interface{}{1, nil, 2}),
			output: false,
		},
		{
			name:   "Example 3",
			p:      CreateTreeFromArray([]interface{}{1, 2, 1}),
			q:      CreateTreeFromArray([]interface{}{1, 1, 2}),
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, isSameTree(testCase.p, testCase.q), testCase.output)
		})
	}
}
