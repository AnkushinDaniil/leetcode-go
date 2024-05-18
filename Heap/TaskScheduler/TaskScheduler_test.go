package taskScheduler

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		tasks  []byte
		n      int
		output int
	}{
		{
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'C', 'C'},
			n:      1,
			output: 7,
		},
		{
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:      2,
			output: 8,
		},
		{
			tasks:  []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:      1,
			output: 6,
		},
		{
			tasks:  []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:      3,
			output: 10,
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.Equal(t, testCase.output, leastInterval(testCase.tasks, testCase.n))
		})
	}
}
