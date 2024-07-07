package palindromeNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	name   string
	n      uint32
	output uint32
}

var tests = []test{
	{
		name:   "00000010100101000001111010011100",
		n:      43261596,
		output: 964176192,
	},
	{
		name:   "11111111111111111111111111111101",
		n:      4294967293,
		output: 3221225471,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.output, reverseBits(testCase.n))
		})
	}
}
