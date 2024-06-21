package mergeTripletsToFormTargetTriplet

import (
	"sync"
	"sync/atomic"
)

func mergeTriplets(triplets [][]int, target []int) bool {
	merged := [3]bool{false, false, false}
	var counter uint8 = 0

	isOk := func(i int) bool {
		ok := true
		for j := 0; j < len(triplets[i]); j++ {
			ok = ok && (triplets[i][j] <= target[j])
		}
		return ok
	}

	for i := 0; i < len(triplets); i++ {
		if isOk(i) {
			for j := 0; j < 3; j++ {
				if !merged[j] && triplets[i][j] == target[j] {
					merged[j] = true
					counter++
					if counter == 3 {
						return true
					}
				}
			}
		}
	}

	return counter == 3
}

func mergeTripletsConcur(triplets [][]int, target []int) bool {
	var merged [3]atomic.Bool

	isOk := func(i int) bool {
		var counter atomic.Int32
		counter.Store(0)
		wg := &sync.WaitGroup{}
		for j := 0; j < len(triplets[i]); j++ {
			wg.Add(1)
			go func(jj int) {
				if triplets[i][jj] <= target[jj] {
					counter.Add(1)
				}
				wg.Done()
			}(j)
		}
		wg.Wait()
		return counter.Load() == 3
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < len(triplets); i++ {
		wg.Add(1)
		go func(ii int) {
			if isOk(ii) {
				for j := 0; j < 3; j++ {
					wg.Add(1)
					go func(jj int) {
						if triplets[ii][jj] == target[jj] {
							merged[jj].Store(true)
						}
						wg.Done()
					}(j)
				}
			}
			wg.Done()
		}(i)
	}

	var res atomic.Bool
	res.Store(true)

	wg.Wait()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(j int) {
			res.Store(res.Load() && merged[j].Load())
			wg.Done()
		}(i)
	}
	wg.Wait()
	return res.Load()
}

func mergeTripletsOld(triplets [][]int, target []int) bool {
	merge := func(i, j int) {
		for k := 0; k < len(triplets[j]); k++ {
			triplets[j][k] = max(triplets[i][k], triplets[j][k])
		}
	}

	isOk := func(i int) bool {
		ok := true
		for j := 0; j < len(triplets[i]); j++ {
			ok = ok && (triplets[i][j] <= target[j])
		}
		return ok
	}

	i := 0
	for ; i < len(triplets); i++ {
		if isOk(i) {
			break
		}
	}
	if i == len(triplets) {
		return false
	}

	for j := i + 1; j < len(triplets); j++ {
		if isOk(j) {
			merge(i, j)
			i = j
		}
	}

	isEqual := true
	for j := 0; j < len(triplets[i]); j++ {
		isEqual = isEqual && triplets[i][j] == target[j]
	}

	return isEqual
}
