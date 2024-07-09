package convertSortedArrayToBinarySearchTree

import (
	"strconv"
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

type testCase struct {
	nums   []int
	output *TreeNode
}

var testTable = []testCase{
	{
		nums:   []int{-10, -3, 0, 5, 9},
		output: CreateTreeFromArray([]interface{}{0, -10, 5, nil, -3, nil, 9}),
	},
	{
		nums:   []int{1, 3},
		output: CreateTreeFromArray([]interface{}{1, nil, 3}),
	},
}

func Test(t *testing.T) {
	for i, testCase := range testTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, sortedArrayToBST(testCase.nums), testCase.output)
		})
	}
}
