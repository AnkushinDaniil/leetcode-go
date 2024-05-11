package letterCombinationsOfAPhoneNumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		digits string
		output []string
	}{
		{
			digits: "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			digits: "",
			output: []string{},
		},
		{
			digits: "2",
			output: []string{"a", "b", "c"},
		},
		{
			digits: "3",
			output: []string{"d", "e", "f"},
		},
		{
			digits: "7",
			output: []string{"p", "q", "r", "s"},
		},
	}

	for i, testCase := range testTable {
		t.Run(fmt.Sprintf("Example №%v", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, letterCombinations(testCase.digits))
		})
	}
}

func BenchmarkI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("4632")
	}
}

func BenchmarkRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsRange("4632")
	}
}

func BenchmarkRangeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsMap("4632")
	}
}

func BenchmarkRangeEmply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsEmpty("4632")
	}
}

func BenchmarkRangeRecurse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsRecurse("4632")
	}
}

// BenchmarkI-10                     645223              1850 ns/op            1736 B/op         83 allocs/op
// BenchmarkRange-10                 331414              3608 ns/op            2072 B/op         83 allocs/op
// BenchmarkRangeMap-10              188852              6170 ns/op            6376 B/op        169 allocs/op
// BenchmarkRangeEmply-10            443802              2625 ns/op            4792 B/op         90 allocs/op
// BenchmarkRangeRecurse-10          225898              5215 ns/op            4932 B/op        128 allocs/op
