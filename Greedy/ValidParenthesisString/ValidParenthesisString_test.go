package validParenthesisString

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	s      string
	output bool
}{
	{
		s:      "()",
		output: true,
	},
	{
		s:      "(*)",
		output: true,
	},
	{
		s:      "(*))",
		output: true,
	},
	{
		s:      "(()",
		output: false,
	},
	{
		s:      "**",
		output: true,
	},
	{
		s:      "(((**",
		output: false,
	},
	{
		s:      "(",
		output: false,
	},
	{
		s:      "*",
		output: true,
	},
	{
		s:      ")*",
		output: false,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].s), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				checkValidString(testTable[i].s),
			)
		})
	}
}
