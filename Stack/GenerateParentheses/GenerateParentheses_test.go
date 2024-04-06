package generateParentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		n      int
		output []string
	}{
		{
			name:   "Example 1",
			n:      3,
			output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name:   "Example 2",
			n:      1,
			output: []string{"()"},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, generateParenthesis(testCase.n), testCase.output)
		})
	}
}
