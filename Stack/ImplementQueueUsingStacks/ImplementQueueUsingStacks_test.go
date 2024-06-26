package implementQueueUsingStacks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	functions []string
	input     []int
	output    []interface{}
}

type Queue interface {
	Push(int)
	Pop() int
	Peek() int
	Empty() bool
}

var tests = []test{
	{
		functions: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
		input:     []int{0, 1, 2, 0, 0, 0},
		output:    []interface{}{nil, nil, nil, 1, 1, false},
	},
	{
		functions: []string{"MyQueue", "empty"},
		input:     []int{0, 0},
		output:    []interface{}{nil, true},
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		t.Run(fmt.Sprint(testCase.functions), func(t *testing.T) {
			runTest(t, testCase)
		})
	}
}

func runTest(t *testing.T, testCase test) {
	var queue MyQueue
	output := make([]interface{}, len(testCase.functions))
	for i, function := range testCase.functions {
		switch function {
		case "MyQueue":
			queue = Constructor()
		case "push":
			queue.Push(testCase.input[i])
		case "pop":
			output[i] = queue.Pop()
		case "peek":
			output[i] = queue.Peek()
		case "empty":
			output[i] = queue.Empty()
		}
	}
	assert.Equal(t, testCase.output, output)
}

func run(testCase test, queue Queue) {
	output := make([]interface{}, len(testCase.functions))
	for i, function := range testCase.functions {
		switch function {
		case "push":
			queue.Push(testCase.input[i])
		case "pop":
			output[i] = queue.Pop()
		case "peek":
			output[i] = queue.Peek()
		case "empty":
			output[i] = queue.Empty()
		}
	}
}

func BenchmarkMy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := Constructor()
		run(tests[0], &q)
	}
}

func BenchmarkLCT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q MyQueueLCT
		run(tests[0], &q)
	}
}

// BenchmarkMy-10           8106213               139.8 ns/op           176 B/op              4 allocs/op
// BenchmarkLCT-10          7344652               165.8 ns/op           152 B/op              5 allocs/op
