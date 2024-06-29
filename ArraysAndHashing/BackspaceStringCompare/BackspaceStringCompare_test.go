package romanToInteger

import "testing"

type test struct {
	s, t   string
	output bool
}

var tests = []test{
	{
		s:      "bxj##tw",
		t:      "bxj###tw",
		output: false,
	},
	{
		s:      "ab#c",
		t:      "ad#c",
		output: true,
	},
	{
		s:      "ab##",
		t:      "c#d#",
		output: true,
	},
	{
		s:      "a#c",
		t:      "b",
		output: false,
	},
	{
		s:      "bxj##tw",
		t:      "bxo#j##tw",
		output: true,
	},
	{
		s:      "nzp#o#g",
		t:      "b#nzp#o#g",
		output: true,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := backspaceCompare(testCase.s, testCase.t)
	if testCase.output != res {
		t.Errorf("test:%v, \nexpected: %v, \ngot: %v", testCase.s, testCase.output, res)
	}
}

func Benchmark_backspaceCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompare(
			"######fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f##",
			"######fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f##",
		)
	}
}

func Benchmark_backspaceCompareOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		backspaceCompareOld(
			"######fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f##",
			"######fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f##",
		)
	}
}

// "aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#aa#"
// Benchmark_backspaceCompare-10              81054             14056 ns/op          0 B/op           0 allocs/op
// Benchmark_backspaceCompareOld-10          108727             10790 ns/op      12544 B/op           4 allocs/op

//"######fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f########fffff###ffff####f##"
// Benchmark_backspaceCompare-10             495176              2399 ns/op          0 B/op           0 allocs/op
// Benchmark_backspaceCompareOld-10          317433              3681 ns/op       2816 B/op           2 allocs/op
