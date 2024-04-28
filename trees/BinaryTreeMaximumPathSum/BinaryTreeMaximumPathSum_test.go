package binaryTreeMaximumPathSum

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
			root:   CreateTreeFromArray([]interface{}{1, 2, 3}),
			output: 6,
		},
		{
			name:   "Example 2",
			root:   CreateTreeFromArray([]interface{}{-10, 9, 20, nil, nil, 15, 7}),
			output: 42,
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{-3}),
			output: -3,
		},
		{
			name:   "Example 3",
			root:   CreateTreeFromArray([]interface{}{1, -2, 3}),
			output: 4,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				maxPathSum(testCase.root),
			)
		})
	}
}

func Benchmark(b *testing.B) {
	root := CreateTreeFromArray([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	for i := 0; i < b.N; i++ {
		maxPathSum(root)
	}
}

func BenchmarkMax(b *testing.B) {
	root := CreateTreeFromArray([]interface{}{-10, 9, 20, nil, nil, 15, 7})
	for i := 0; i < b.N; i++ {
		maxPathSum_(root)
	}
}
