package fibonaccinumber

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
  tests := []struct{
    input, output int
  }{
    {2, 1},
    {3, 2},
    {4, 3},
  }

  for _, test := range tests {
    t.Run(strconv.Itoa(test.input), func(t *testing.T) {
      assert.Equal(t, test.output, fib(test.input))
    })
  }
}

func BenchmarkFibN(  b *testing.B,) {
  for i := 0; i < b.N; i++ {
    fibN(1000)
  }
}

func BenchmarkFibLogN(  b *testing.B,) {
  for i := 0; i < b.N; i++ {
    fib(1000)
  }
}
