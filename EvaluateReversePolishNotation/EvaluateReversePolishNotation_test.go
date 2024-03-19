package evaluateReversePolishNotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		tokens []string
		output int
	}{
		{
			name:   "Example 1",
			tokens: []string{"2", "1", "+", "3", "*"},
			output: 9,
		},
		{
			name:   "Example 2",
			tokens: []string{"4", "13", "5", "/", "+"},
			output: 6,
		},
		{
			name:   "Example 3",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			output: 22,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, evalRPN(testCase.tokens), testCase.output)
		})
	}
}
