package decodeWays

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
		s:      "12",
		output: 2,
	},
	{
		s:      "226",
		output: 3,
	},
	{
		s:      "0",
		output: 0,
	},
	{
		s:      "27",
		output: 1,
	},
	{
		s:      "100",
		output: 0,
	},
	{
		s:      "2101",
		output: 1,
	},
	{
		s:      "12",
		output: 2,
	},
	{
		s:      "230",
		output: 0,
	},
	{
		s:      "10",
		output: 1,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.s), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				numDecodings(testCase.s),
			)
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodings("1171615262627178281272")
	}
}

func BenchmarkMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodingsMy("1171615262627178281272")
	}
}

func BenchmarkLeetcodeTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodingsLeetcodeTime("1171615262627178281272")
	}
}

func BenchmarkLeetcodeMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numDecodingsLeetcodeMemory("1171615262627178281272")
	}
}

// Benchmark-10                    94698210                10.83 ns/op            0 B/op          0 allocs/op
// BenchmarkMy-10                  43289198                27.16 ns/op            0 B/op          0 allocs/op
// BenchmarkLeetcodeTime-10          927482                1209 ns/op             965 B/op        3 allocs/op
// BenchmarkLeetcodeMemory-10      20973734                56.54 ns/op            0 B/op          0 allocs/op
