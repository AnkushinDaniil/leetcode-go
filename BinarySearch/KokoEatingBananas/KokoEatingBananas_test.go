package kokoEatingBananas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		piles  []int
		h      int
		output int
	}{
		{
			name:   "Example 1",
			piles:  []int{3, 6, 7, 11},
			h:      8,
			output: 4,
		},
		{
			name:   "Example 2",
			piles:  []int{30, 11, 23, 4, 20},
			h:      5,
			output: 30,
		},
		{
			name:   "Example 3",
			piles:  []int{30, 11, 23, 4, 20},
			h:      6,
			output: 23,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, minEatingSpeed(testCase.piles, testCase.h), testCase.output)
		})
	}
}
