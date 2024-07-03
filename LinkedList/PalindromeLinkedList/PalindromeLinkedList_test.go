package palindromeLinkedList

import (
	"testing"

	"github.com/AnkushinDaniil/leetcode-go/linkedList"
	"github.com/stretchr/testify/assert"
)

type test struct {
	head   *ListNode
	output bool
}

var tests = []test{
	{
		head:   linkedList.CreateListFromArray([]int{1, 2, 3, 2, 1}),
		output: true,
	},
	{
		head:   linkedList.CreateListFromArray([]int{1, 2, 2, 1}),
		output: true,
	},
	{
		head:   linkedList.CreateListFromArray([]int{1, 2}),
		output: false,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := isPalindrome(testCase.head)
	assert.Equal(t, testCase.output, res)
}
