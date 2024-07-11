package kClosestPointstoOrigin

type point struct {
	coords         []int
	distanceSquare int
}

func sumOfSquares(x, y int) int {
	return x*x + y*y
}

func NewPoint(p []int) *point {
	return &point{
		coords:         p,
		distanceSquare: sumOfSquares(p[0], p[1]),
	}
}

type heapPoint []*point

func (h heapPoint) Compare(i, j int) bool {
	return h[i].distanceSquare >= h[j].distanceSquare
}

func (h heapPoint) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func NewHeap(capacity int) *heapPoint {
	heap := make(heapPoint, 0, capacity)
	return &heap
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (h *heapPoint) heapifyUp(i int) {
	for i > 0 {
		p := parent(i)
		if h.Compare(i, p) {
			h.Swap(i, p)
		} else {
			return
		}
		i = p
	}
}

func (h *heapPoint) heapifyDown(i int) {
	for {
		highest := i

		l := left(i)
		if l < len(*h) && h.Compare(l, i) {
			highest = l
		}

		r := right(i)
		if r < len(*h) && h.Compare(r, highest) {
			highest = r
		}

		if i == highest {
			return
		}
		h.Swap(i, highest)
		i = highest
	}
}

func (h *heapPoint) Insert(p *point) {
	*h = append(*h, p)
	h.heapifyUp(len(*h) - 1)
	// fmt.Println("Insert: ", h.isHeap(0))
}

func (h *heapPoint) Pop() *point {
	res := (*h)[0]
	h.Swap(0, len(*h)-1)
	*h = (*h)[:len(*h)-1]
	h.heapifyDown(0)
	// fmt.Println("Pop: ", h.isHeap(0))
	return res
}

// func (h *heapPoint) isHeap(i int) bool {
// 	if i >= len(*h) {
// 		return true
// 	}
// 	return h.Compare(parent(i), i) && h.isHeap(left(i)) && h.isHeap(right(i))
// }

func kClosest(points [][]int, k int) [][]int {
	if k >= len(points) {
		return points
	}
	heap := NewHeap(k + 1)
	i := 0
	for ; i < k; i++ {
		heap.Insert(NewPoint(points[i]))
	}
	for ; i < len(points); i++ {
		heap.Insert(NewPoint(points[i]))
		heap.Pop()
	}
	i = 0
	for ; len(*heap) > 0; i++ {
		points[i] = heap.Pop().coords
	}
	return points[:k]
}
