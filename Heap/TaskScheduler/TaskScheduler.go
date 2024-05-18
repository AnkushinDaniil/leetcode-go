package taskScheduler

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

func leastIntervalHeap(tasks []byte, n int) int {
	table := [26]int{}
	for i := 0; i < len(tasks); i++ {
		table[tasks[i]-'A']++
	}

	heap := NewMaxHeap(27)
	for i := 0; i < 26; i++ {
		if table[i] > 0 {
			heap.Insert(table[i])
		}
	}

	queue := make([][2]int, 0, len(heap)+n)
	time := 0
	num := 0
	for ; len(heap) > 0 || len(queue) > 0; time++ {
		if len(heap) > 0 {
			num = heap.ExtractMax()
			if num > 1 {
				queue = append(queue, [2]int{num - 1, time + n})
			}
		}
		if len(queue) > 0 && queue[0][1] <= time {
			heap.Insert(queue[0][0])
			queue = queue[1:]
		}
	}

	return time
}

// -----------------------------------------

func leastInterval(tasks []byte, n int) int {
	table := [26]int{}
	for i := 0; i < len(tasks); i++ {
		table[tasks[i]-'A']++
	}

	maxCount, maxCountRepeats := table[0], 1
	for i := 1; i < 26; i++ {
		if table[i] > maxCount {
			maxCount = table[i]
			maxCountRepeats = 1
		} else if table[i] == maxCount {
			maxCountRepeats++
		}
	}

	res := (n+1)*(maxCount-1) + maxCountRepeats
	if res > len(tasks) {
		return res
	} else {
		return len(tasks)
	}
}
