package palindromicSubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	s      string
	output int
}{
	{
		s:      "abc",
		output: 3,
	},
	{
		s:      "aaa",
		output: 6,
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				countSubstrings(testCase.s),
			)
		})
	}
}
