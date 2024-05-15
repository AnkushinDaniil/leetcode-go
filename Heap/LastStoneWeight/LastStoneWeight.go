package lastStoneWeight

type MaxHeap []int

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func NewMaxHeap(cap int) MaxHeap {
	res := make([]int, 0, cap)
	return MaxHeap(res)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Insert(k int) {
	i := len((*h))
	(*h) = append((*h), k)

	for i != 0 && (*h)[parent(i)] < (*h)[i] {
		(*h).Swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MaxHeap) Heapify(i int) {
	l := left(i)
	r := right(i)
	biggest := i
	if l < len((*h)) && (*h)[l] > (*h)[i] {
		biggest = l
	}
	if r < len((*h)) && (*h)[r] > (*h)[biggest] {
		biggest = r
	}
	if biggest != i {
		h.Swap(i, biggest)
		h.Heapify(biggest)
	}
}

func (h *MaxHeap) ExtractMax() int {
	res := (*h)[0]
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
	return res
}

func lastStoneWeight(stones []int) int {
	heap := NewMaxHeap(len(stones))
	for i := 0; i < len(stones); i++ {
		heap.Insert(stones[i])
	}
	for len(heap) > 1 {
		first := heap.ExtractMax()
		second := heap.ExtractMax()
		if second < first {
			heap.Insert(first - second)
		}
	}
	if len(heap) == 0 {
		return 0
	}
	return heap[0]
}
