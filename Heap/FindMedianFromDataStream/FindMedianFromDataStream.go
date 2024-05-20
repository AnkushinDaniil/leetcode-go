package findMedianFromDataStream

const HEAP_CAPACITY = 2

func parent(i int) int {
	return (i - 1) >> 1
}

func left(i int) int {
	return i<<1 + 1
}

func right(i int) int {
	return i<<1 + 2
}

type MinHeap []int

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
	var smallest, l, r int

	for {
		l = left(i)
		r = right(i)
		smallest = i
		if l < len((*h)) && (*h)[l] < (*h)[i] {
			smallest = l
		}
		if r < len((*h)) && (*h)[r] < (*h)[smallest] {
			smallest = r
		}
		if smallest != i {
			h.Swap(i, smallest)
			i = smallest
		} else {
			return
		}
	}
}

func (h *MinHeap) ExtractMin() int {
	res := (*h)[0]
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
	return res
}

func (h *MinHeap) GetMin() int {
	return (*h)[0]
}

type MaxHeap []int

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
	var biggest, l, r int

	for {
		l = left(i)
		r = right(i)
		biggest = i
		if l < len((*h)) && (*h)[l] > (*h)[i] {
			biggest = l
		}
		if r < len((*h)) && (*h)[r] > (*h)[biggest] {
			biggest = r
		}
		if biggest != i {
			h.Swap(i, biggest)
			i = biggest
		} else {
			return
		}
	}
}

func (h *MaxHeap) ExtractMax() int {
	res := (*h)[0]
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
	return res
}

func (h *MaxHeap) GetMax() int {
	return (*h)[0]
}

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: NewMaxHeap(HEAP_CAPACITY),
		minHeap: NewMinHeap(HEAP_CAPACITY),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if (len(mf.maxHeap)+len(mf.minHeap))&1 == 0 {
		mf.minHeap.Insert(num)
		mf.maxHeap.Insert(mf.minHeap.ExtractMin())
	} else {
		mf.maxHeap.Insert(num)
		mf.minHeap.Insert(mf.maxHeap.ExtractMax())
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if (len(mf.maxHeap)+len(mf.minHeap))&1 == 0 {
		return float64(mf.maxHeap.GetMax()+mf.minHeap.GetMin()) / 2.0
	} else {
		return float64(mf.maxHeap.GetMax())
	}
}
