package gasStation

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	gas    []int
	cost   []int
	output int
}{
	{
		gas:    []int{1, 2, 3, 4, 5},
		cost:   []int{3, 4, 5, 1, 2},
		output: 3,
	},
	{
		gas:    []int{2, 3, 4},
		cost:   []int{3, 4, 3},
		output: -1,
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].gas), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				canCompleteCircuit(testTable[i].gas, testTable[i].cost),
			)
		})
	}
}

func Benchmark_canCompleteCircuit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuit(testTable[0].gas, testTable[0].cost)
	}
}

func Benchmark_canCompleteCircuitLCT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuitLCT(testTable[0].gas, testTable[0].cost)
	}
}

func Benchmark_canCompleteCircuitLCM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		canCompleteCircuitLCM(testTable[0].gas, testTable[0].cost)
	}
}

// Benchmark_canCompleteCircuit-10         313765616                3.909 ns/op             0 B/op          0 allocs/op
// Benchmark_canCompleteCircuitLCT-10      309766718                4.069 ns/op             0 B/op          0 allocs/op
// Benchmark_canCompleteCircuitLCM-10      80808308                14.87 ns/op              0 B/op          0 allocs/op
