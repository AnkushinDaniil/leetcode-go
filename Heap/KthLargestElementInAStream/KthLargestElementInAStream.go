package kthLargestElementInAStream

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

// func (h *MinHeap) Decrease(i, newValue int) {
// 	(*h)[i] = newValue
// 	for i != 0 && (*h)[parent(i)] > (*h)[i] {
// 		h.Swap(i, parent(i))
// 		i = parent(i)
// 	}
// }

func (h *MinHeap) Heapify(i int) {
	l := left(i)
	r := right(i)
	smallest := i
	if l < len((*h)) && (*h)[l] < (*h)[i] {
		smallest = l
	}
	if r < len((*h)) && (*h)[r] < (*h)[smallest] {
		smallest = r
	}
	if smallest != i {
		h.Swap(i, smallest)
		h.Heapify(smallest)
	}
}

func (h *MinHeap) ExtractMin() {
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
}

// func (h *MinHeap) Delete(i int) {
// 	h.Decrease(i, math.MinInt)
// 	h.ExtractMin()
// }

func (h *MinHeap) GetMin() int {
	return (*h)[0]
}

type KthLargest struct {
	heap *MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	heap := NewMinHeap(k + 1)
	for i := 0; i < len(nums); i++ {
		heap.Insert(nums[i])
		if len(heap) > k {
			heap.ExtractMin()
		}
	}
	return KthLargest{
		heap: &heap,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	this.heap.Insert(val)
	if len(*this.heap) > this.k {
		this.heap.ExtractMin()
	}
	return this.heap.GetMin()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
