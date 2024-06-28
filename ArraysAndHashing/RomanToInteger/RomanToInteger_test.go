package romanToInteger

import "testing"

type test struct {
	s      string
	output int
}

var tests = []test{
	{
		s:      "III",
		output: 3,
	},
	{
		s:      "LVIII",
		output: 58,
	},
	{
		s:      "MCMXCIV",
		output: 1994,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := romanToInt(testCase.s)
	if testCase.output != res {
		t.Errorf("expected: %v, \ngot: %v", testCase.output, res)
	}
}
