package palindromeNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name   string
	n      int
	output bool
}

var tests = []test{
	// {
	// 	name:   "OK",
	// 	n:      121,
	// 	output: true,
	// },
	// {
	// 	name:   "negative",
	// 	n:      -121,
	// 	output: false,
	// },
	// {
	// 	name:   "not a palindrome",
	// 	n:      10,
	// 	output: false,
	// },
	// {
	// 	name:   "zero",
	// 	n:      0,
	// 	output: true,
	// },
	// {
	// 	name:   "-1",
	// 	n:      -1,
	// 	output: false,
	// },
	{
		name:   "11",
		n:      11,
		output: true,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.output, isPalindrome(testCase.n))
		})
	}
}
