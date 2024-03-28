package timeBasedKeyValueStore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	arg1 string
	arg2 string
	arg3 int
}

func Test(t *testing.T) {
	testTable := []struct {
		name      string
		functions []string
		input     []Input
		output    [][]string
	}{
		{
			name: "Example 1",
			functions: []string{
				"TimeMap",
				"set",
				"get",
				"get",
				"set",
				"get",
				"get",
			},
			input: []Input{
				{},
				{arg1: "foo", arg2: "bar", arg3: 1},
				{arg1: "foo", arg3: 1},
				{arg1: "foo", arg3: 3},
				{arg1: "foo", arg2: "bar2", arg3: 4},
				{arg1: "foo", arg3: 4},
				{arg1: "foo", arg3: 5},
			},
			output: [][]string{{}, {}, {"bar"}, {"bar"}, {}, {"bar2"}, {"bar2"}},
		},
		{
			name: "Example 2",
			functions: []string{
				"TimeMap", "set", "set", "get", "get", "get", "get", "get",
			},
			input: []Input{
				{},
				{arg1: "love", arg2: "high", arg3: 10},
				{arg1: "love", arg2: "low", arg3: 20},
				{arg1: "love", arg3: 5},
				{arg1: "love", arg3: 10},
				{arg1: "love", arg3: 15},
				{arg1: "love", arg3: 20},
				{arg1: "love", arg3: 25},
			},
			output: [][]string{{}, {}, {}, {""}, {"high"}, {"high"}, {"low"}, {"low"}},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			var timeMap TimeMap
			res := make([][]string, 0, len(testCase.functions))
			for i := 0; i < len(testCase.functions); i++ {
				switch testCase.functions[i] {
				case "TimeMap":
					timeMap = Constructor()
					res = append(res, []string{})
				case "set":
					timeMap.Set(
						testCase.input[i].arg1,
						testCase.input[i].arg2,
						testCase.input[i].arg3,
					)
					res = append(res, []string{})
				case "get":
					res = append(
						res,
						[]string{timeMap.Get(
							testCase.input[i].arg1,
							testCase.input[i].arg3,
						)},
					)
				}
			}
			assert.Equal(t, res, testCase.output)
		})
	}
}
