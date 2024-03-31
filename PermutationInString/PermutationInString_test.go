package permutationInString

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name   string
		s1     string
		s2     string
		output bool
	}{
		{
			name:   "Example 1",
			s1:     "ab",
			s2:     "eidbaooo",
			output: true,
		},
		{
			name:   "Example 2",
			s1:     "ab",
			s2:     "eidboaoo",
			output: false,
		},
		{
			name:   "Example 3",
			s1:     "adc",
			s2:     "dcda",
			output: true,
		},
		{
			name:   "Example 4",
			s1:     "ab",
			s2:     "a",
			output: false,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, checkInclusion(testCase.s1, testCase.s2), testCase.output)
		})
	}
}
