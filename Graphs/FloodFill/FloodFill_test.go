package floodFill

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	image, output [][]int
	sr, sc, color int
}{
	{
		image:  [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr:     1,
		sc:     1,
		color:  2,
		output: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		image:  [][]int{{0, 0, 0}, {0, 0, 0}},
		sr:     0,
		sc:     0,
		color:  0,
		output: [][]int{{0, 0, 0}, {0, 0, 0}},
	},
}

func Test(t *testing.T) {
	for i := 0; i < len(tests); i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(
				t,
				tests[i].output,
				floodFillQueue(tests[i].image, tests[i].sr, tests[i].sc, tests[i].color),
			)
		})
	}
}

func Benchmark_floodFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floodFill(tests[0].image, tests[0].sr, tests[0].sc, tests[0].color)
	}
}

func Benchmark_floodFillQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		floodFillQueue(tests[0].image, tests[0].sr, tests[0].sc, tests[0].color)
	}
}
