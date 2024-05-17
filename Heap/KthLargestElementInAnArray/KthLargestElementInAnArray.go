package kthLargestElementInAnArray

type MinHeap []int

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func NewMinHeap(cap int) MinHeap {
	res := make([]int, 0, cap)
	return MinHeap(res)
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Insert(k int) {
	i := len((*h))
	(*h) = append((*h), k)

	for i != 0 && (*h)[parent(i)] > (*h)[i] {
		(*h).Swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MinHeap) Heapify(i int) {
	l := left(i)
	r := right(i)
	biggest := i
	if l < len((*h)) && (*h)[l] < (*h)[i] {
		biggest = l
	}
	if r < len((*h)) && (*h)[r] < (*h)[biggest] {
		biggest = r
	}
	if biggest != i {
		h.Swap(i, biggest)
		h.Heapify(biggest)
	}
}

func (h *MinHeap) ExtractMin() {
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
}

func (h *MinHeap) GetMin() int {
	return (*h)[0]
}

type KthLargest struct {
	heap *MinHeap
	k    int
}

func findKthLargest(nums []int, k int) int {
	heap := NewMinHeap(k + 1)
	i := 0
	for ; i < k; i++ {
		heap.Insert(nums[i])
	}
	for ; i < len(nums); i++ {
		heap.Insert(nums[i])
		heap.ExtractMin()
	}
	return heap.GetMin()
}
