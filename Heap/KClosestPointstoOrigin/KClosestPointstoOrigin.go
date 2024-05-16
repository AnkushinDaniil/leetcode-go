package kClosestPointstoOrigin

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

var (
	res   [][]int
	point []int = make([]int, 2)
)

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	res[i], res[j] = res[j], res[i]
}

func (h *MaxHeap) Insert(k int) {
	i := len((*h))
	(*h) = append((*h), k)
	res = append(res, point)

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

func (h *MaxHeap) ExtractMax() {
	(*h)[0] = (*h)[len((*h))-1]
	res[0] = res[len(res)-1]
	(*h) = (*h)[:len((*h))-1]
	res = res[:len(res)-1]
	(*h).Heapify(0)
}

func (h *MaxHeap) GetMax() int {
	return (*h)[0]
}

func kClosest(points [][]int, k int) [][]int {
	res = make([][]int, 0, len(points))
	heap := NewMaxHeap(len(points))
	i := 0
	for ; i < k; i++ {
		point = points[i]
		heap.Insert(point[0]*point[0] + point[1]*point[1])
	}
	for ; i < len(points); i++ {
		point = points[i]
		heap.Insert(point[0]*point[0] + point[1]*point[1])
		heap.ExtractMax()
	}
	return res
}
