package сloneGraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		node *Node
	}{
		{
			node: Create([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			node: Create([][]int{{}}),
		},
		{
			node: Create([][]int{}),
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			assert.EqualValues(t, testCase.node, cloneGraph(testCase.node))
		})
	}
}

func Create(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	nodes := make([]Node, len(adjList))
	for i := 0; i < len(adjList); i++ {
		nodes[i].Val = i + 1
		nodes[i].Neighbors = make([]*Node, len(adjList[i]))
		for j := 0; j < len(adjList[i]); j++ {
			nodes[i].Neighbors[j] = &nodes[adjList[i][j]-1]
		}
	}
	return &nodes[0]
}

func BenchmarkDeque(b *testing.B) {
	graph := Create([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	for i := 0; i < b.N; i++ {
		cloneGraphDeque(graph)
	}
}

func BenchmarkQueue(b *testing.B) {
	graph := Create([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	for i := 0; i < b.N; i++ {
		cloneGraphQueue(graph)
	}
}

func BenchmarkRecursion(b *testing.B) {
	graph := Create([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}})
	for i := 0; i < b.N; i++ {
		cloneGraph(graph)
	}
}

// BenchmarkDeque-10                1145928              1068 ns/op            2904 B/op             10 allocs/op
// BenchmarkQueue-10                1000000              1075 ns/op            2904 B/op             10 allocs/op
// BenchmarkRecursion-10            1461088               820.4 ns/op          2904 B/op             10 allocs/op
