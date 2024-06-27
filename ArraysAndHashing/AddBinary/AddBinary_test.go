package addBinary

import "testing"

type test struct {
	a, b, output string
}

var tests = []test{
	{
		a:      "11",
		b:      "1",
		output: "100",
	},
	{
		a:      "1010",
		b:      "1011",
		output: "10101",
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := addBinary(testCase.a, testCase.b)
	if testCase.output != res {
		t.Errorf("expected: %v, \ngot: %v", testCase.output, res)
	}
}
