package dailyTemperatures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		name         string
		temperatures []int
		output       []int
	}{
		{
			name:         "Example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			output:       []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example 2",
			temperatures: []int{30, 40, 50, 60},
			output:       []int{1, 1, 1, 0},
		},
		{
			name:         "Example 3",
			temperatures: []int{30, 60, 90},
			output:       []int{1, 1, 0},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, dailyTemperatures(testCase.temperatures), testCase.output)
		})
	}
}
