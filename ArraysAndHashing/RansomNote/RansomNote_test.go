package ransomNote

import "testing"

type test struct {
	ransomNote, magazine string
	output               bool
}

var tests = []test{
	{
		ransomNote: "a",
		magazine:   "b",
		output:     false,
	},
	{
		ransomNote: "aa",
		magazine:   "ab",
		output:     false,
	},
	{
		ransomNote: "aa",
		magazine:   "aab",
		output:     true,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := canConstruct(testCase.ransomNote, testCase.magazine)
	if testCase.output != res {
		t.Errorf("expected: %v, \ngot: %v", testCase.output, res)
	}
}
