package firstBadVersion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n, bad int
}

var tests = []testCase{
	{
		n:   5,
		bad: 4,
	},
	{
		n:   1,
		bad: 1,
	},
}

func Test(t *testing.T) {
	for _, test := range tests {
		runTest(t, test)
	}
}

func runTest(t *testing.T, test testCase) {
	isBadVersion = func(x int) bool {
		return x >= test.bad
	}
	assert.Equal(t, test.bad, firstBadVersion(test.n))
}
