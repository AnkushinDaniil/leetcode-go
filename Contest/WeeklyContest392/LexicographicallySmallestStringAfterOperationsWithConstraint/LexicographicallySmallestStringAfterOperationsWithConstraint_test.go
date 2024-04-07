package binarySearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		s      string
		k      int
		output string
	}{
		{
			name:   "Example 1",
			s:      "zbbz",
			k:      3,
			output: "aaaz",
		},
		{
			name:   "Example 2",
			s:      "xaxcd",
			k:      4,
			output: "aawcd",
		},
		{
			name:   "Example 3",
			s:      "lol",
			k:      0,
			output: "lol",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, getSmallestString(testCase.s, testCase.k), testCase.output)
		})
	}
}
