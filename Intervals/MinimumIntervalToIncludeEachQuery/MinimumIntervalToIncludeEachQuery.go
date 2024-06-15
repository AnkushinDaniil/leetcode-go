package minimumIntervalToIncludeEachQuery

import (
	"container/heap"
	"slices"
)

type DefaultNode struct {
	val, idx int
}

type DefaultHeap []DefaultNode

func (h DefaultHeap) Len() int           { return len(h) }
func (h DefaultHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h DefaultHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *DefaultHeap) Push(x any)        { *h = append(*h, x.(DefaultNode)) }

func (h *DefaultHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minIntervalDefault(intervals [][]int, queries []int) []int {
	h := make(DefaultHeap, 0, len(intervals))
	queriesCopy := make([]DefaultNode, len(queries))
	for i := range queries {
		queriesCopy[i] = DefaultNode{
			val: queries[i],
			idx: i,
		}
	}
	slices.SortFunc(queriesCopy, func(a, b DefaultNode) int {
		return a.val - b.val
	})
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	j := 0
	for i := range queriesCopy {
		for ; j < len(intervals) && intervals[j][0] <= queriesCopy[i].val; j++ {
			heap.Push(&h, DefaultNode{
				val: intervals[j][1] - intervals[j][0] + 1,
				idx: intervals[j][1],
			})
		}
		for len(h) > 0 && h[0].idx < queriesCopy[i].val {
			heap.Pop(&h)
		}
		if len(h) == 0 {
			queries[queriesCopy[i].idx] = -1
		} else {
			queries[queriesCopy[i].idx] = h[0].val
		}
	}

	return queries
}
