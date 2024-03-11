package topKFrequentElements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	table := make(map[int]int)

	for _, num := range nums {
		table[num]++
	}

	type Pair struct {
		Key   int
		Value int
	}

	pairs := make([]Pair, len(table))
	i := 0

	for key, val := range table {
		pairs[i] = Pair{
			Key:   key,
			Value: val,
		}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	res := make([]int, k)

	for i := 0; i < k; i++ {
		res[i] = pairs[i].Key
	}

	return res
}
