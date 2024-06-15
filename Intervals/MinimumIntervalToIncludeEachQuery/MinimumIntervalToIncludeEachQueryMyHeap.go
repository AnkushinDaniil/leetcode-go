package minimumIntervalToIncludeEachQuery

import (
	"slices"
)

type MyNode struct {
	val, idx int
}

type MyHeap []MyNode

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func NewMyHeap(cap int) MyHeap {
	res := make([]MyNode, 0, cap)
	return MyHeap(res)
}

func (h *MyHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MyHeap) Insert(k MyNode) {
	i := len((*h))
	(*h) = append((*h), k)

	for i != 0 && (*h)[parent(i)].val > (*h)[i].val {
		(*h).Swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MyHeap) Heapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < len((*h)) && (*h)[l].val < (*h)[i].val {
		smallest = l
	}
	if r < len((*h)) && (*h)[r].val < (*h)[smallest].val {
		smallest = r
	}
	if smallest != i {
		h.Swap(i, smallest)
		h.Heapify(smallest)
	}
}

func (h *MyHeap) ExtractMin() {
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
}

func (h *MyHeap) GetMin() MyNode {
	return (*h)[0]
}

func minInterval(intervals [][]int, queries []int) []int {
	h := make(MyHeap, 0, len(intervals))
	queriesCopy := make([]MyNode, len(queries))
	for i := range queries {
		queriesCopy[i] = MyNode{
			val: queries[i],
			idx: i,
		}
	}
	slices.SortFunc(queriesCopy, func(a, b MyNode) int {
		return a.val - b.val
	})
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	j := 0
	for i := range queriesCopy {
		for ; j < len(intervals) && intervals[j][0] <= queriesCopy[i].val; j++ {
			h.Insert(MyNode{
				val: intervals[j][1] - intervals[j][0] + 1,
				idx: intervals[j][1],
			})
		}
		for len(h) > 0 && h.GetMin().idx < queriesCopy[i].val {
			h.ExtractMin()
		}
		if len(h) == 0 {
			queries[queriesCopy[i].idx] = -1
		} else {
			queries[queriesCopy[i].idx] = h.GetMin().val
		}
	}

	return queries
}
