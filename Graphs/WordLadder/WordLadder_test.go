package wordLadder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTable = []struct {
	beginWord, endWord string
	wordList           []string
	output             int
}{
	{
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
		output:    5,
	},
	{
		beginWord: "hit",
		endWord:   "cog",
		wordList:  []string{"hot", "dot", "dog", "lot", "log"},
		output:    0,
	},
	{
		beginWord: "hot",
		endWord:   "dog",
		wordList:  []string{"hot", "dog"},
		output:    0,
	},
}

func Test(t *testing.T) {
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(
				t,
				testCase.output,
				ladderLengthBytes(testCase.beginWord, testCase.endWord, testCase.wordList),
			)
		})
	}
}

func BenchmarkString(b *testing.B) {
	n := len(testTable)
	for i := 0; i < b.N; i++ {
		ladderLength(testTable[i%n].beginWord, testTable[i%n].endWord, testTable[i%n].wordList)
	}
}

func BenchmarkSB(b *testing.B) {
	n := len(testTable)
	for i := 0; i < b.N; i++ {
		ladderLengthSB(testTable[i%n].beginWord, testTable[i%n].endWord, testTable[i%n].wordList)
	}
}

func BenchmarkBytes(b *testing.B) {
	n := len(testTable)
	for i := 0; i < b.N; i++ {
		ladderLengthBytes(testTable[i%n].beginWord, testTable[i%n].endWord, testTable[i%n].wordList)
	}
}

// BenchmarkString-10        692218              1725 ns/op            1110 B/op         22 allocs/op
// BenchmarkSB-10            556117              2066 ns/op            1214 B/op         29 allocs/op
// BenchmarkBytes-10         413983              2713 ns/op            1219 B/op         33 allocs/op
