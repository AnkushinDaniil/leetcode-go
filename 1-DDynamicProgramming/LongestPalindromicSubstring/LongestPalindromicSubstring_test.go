package longestPalindromicSubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	s      string
	output string
}{
	{
		s:      "bavdqwertyuiopoiuytrewqad",
		output: "qwertyuiopoiuytrewq",
	},
	{
		s:      "babad",
		output: "bab",
	},

	{
		s:      "cbbd",
		output: "bb",
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				longestPalindrome(testCase.s),
			)
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome(testTable[0].s)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestPalindrome2(testTable[0].s)
	}
}

// Benchmark1-10           12695126                91.88 ns/op            0 B/op          0 allocs/op
// Benchmark2-10           12935535                91.56 ns/op            0 B/op          0 allocs/op
