package wordBreak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	s        string
	wordDict []string
	output   bool
}{
	{
		s:        "leetcode",
		wordDict: []string{"leet", "code"},
		output:   true,
	},
	{
		s:        "applepenapple",
		wordDict: []string{"apple", "pen"},
		output:   true,
	},
	{
		s:        "catsandog",
		wordDict: []string{"cats", "dog", "sand", "and", "cat"},
		output:   false,
	},
	{
		s:        "abcd",
		wordDict: []string{"a", "abc", "b", "cd"},
		output:   true,
	},
}

func Test(t *testing.T) {
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test %v", testCase.s), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				wordBreakArray(testCase.s, testCase.wordDict),
			)
		})
	}
}

func Benchmark_wordBreak(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordBreak(
			"applepenappleapplepenappleapplepenapple",
			[]string{"apple", "pen", "applepenapple"},
		)
	}
}

func Benchmark_wordBreakArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordBreakArray(
			"applepenappleapplepenappleapplepenapple",
			[]string{"apple", "pen", "applepenapple"},
		)
	}
}

func Benchmark_wordBreakNeetcode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordBreakNeetcode(
			"applepenappleapplepenappleapplepenapple",
			[]string{"apple", "pen", "applepenapple"},
		)
	}
}

func Benchmark_wordBreakLeetcode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordBreakLeetcode(
			"applepenappleapplepenappleapplepenapple",
			[]string{"apple", "pen", "applepenapple"},
		)
	}
}

// Benchmark_wordBreak-10                   6261477               187.8 ns/op            48 B/op           1 allocs/op
// Benchmark_wordBreakArray-10              7197486               166.5 ns/op             0 B/op           0 allocs/op
// Benchmark_wordBreakNeetcode-10           2088925               573.9 ns/op            48 B/op           1 allocs/op
// Benchmark_wordBreakLeetcode-10            598707              1912 ns/op            3984 B/op          18 allocs/op
