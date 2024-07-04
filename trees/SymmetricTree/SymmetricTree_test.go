package symmetricTree

import (
	"strconv"
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/trees"
	"github.com/stretchr/testify/assert"
)

var CreateTreeFromArray func([]interface{}) *TreeNode = trees.CreateTreeFromArray

type testCase struct {
	root   *TreeNode
	output bool
}

var testTable = []testCase{
	{
		root:   CreateTreeFromArray([]interface{}{1, 2, 2, 3, 4, 4, 3}),
		output: true,
	},
	{
		root:   CreateTreeFromArray([]interface{}{1, 2, 2, nil, 3, nil, 3}),
		output: false,
	},
}

func Test(t *testing.T) {
	for i, testCase := range testTable {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, isSymmetric(testCase.root), testCase.output)
		})
	}
}
