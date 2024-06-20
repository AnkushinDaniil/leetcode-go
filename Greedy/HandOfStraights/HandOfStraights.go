package handOfStraights

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

func isNStraightHand(hand []int, groupSize int) bool {
	minHeap := NewMinHeap(len(hand))
	hashMap := make(map[int]int, len(hand))

	for i := 0; i < len(hand); i++ {
		if _, ok := hashMap[hand[i]]; !ok {
			hashMap[hand[i]] = 1
			minHeap.Insert(hand[i])
		} else {
			hashMap[hand[i]]++
		}
	}

	for len(minHeap) > 0 {
		start := minHeap.GetMin()
		for hashMap[start] > 0 {
			for i := 0; i < groupSize; i++ {
				if hashMap[start+i] == 0 {
					return false
				} else {
					hashMap[start+i]--
				}
			}
		}
		minHeap.ExtractMin()
	}
	return true
}
