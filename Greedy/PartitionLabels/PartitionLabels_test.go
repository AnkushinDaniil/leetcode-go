package partitionLabels

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	s      string
	output []int
}{
	{
		s:      "ababcbacadefegdehijhklij",
		output: []int{9, 7, 8},
	},
	{
		s:      "eccbbbbdec",
		output: []int{10},
	},
}

func Test(t *testing.T) {
	for i := range testTable {
		t.Run(fmt.Sprintf("Test %v", testTable[i].s), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				partitionLabels(testTable[i].s),
			)
		})
	}
}

func Benchmark_partitionLabels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionLabels(
			"erqwrerwqerqwerqwereweqwqrewreqryitutytyuituyituytiuyityutuyityuisfadfdsafsdfdsadfsadfassfadafdghjkghjkghjghjkgjhkgjhghjghjvczxvcxzvcxvzcxzvczxvzcxvczxbnmnbmbnnbmbnmbnmbnmbnm",
		)
	}
}

func Benchmark_partitionLabelsHashMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionLabelsHashMap(
			"erqwrerwqerqwerqwereweqwqrewreqryitutytyuituyituytiuyityutuyityuisfadfdsafsdfdsadfsadfassfadafdghjkghjkghjghjkgjhkgjhghjghjvczxvcxzvcxvzcxzvczxvzcxvczxbnmnbmbnnbmbnmbnmbnmbnm",
		)
	}
}

func Benchmark_partitionLabelsIntervals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		partitionLabelsIntervals(
			"erqwrerwqerqwerqwereweqwqrewreqryitutytyuituyituytiuyityutuyityuisfadfdsafsdfdsadfsadfassfadafdghjkghjkghjghjkgjhkgjhghjghjvczxvcxzvcxvzcxzvczxvzcxvczxbnmnbmbnnbmbnmbnmbnmbnm",
		)
	}
}
