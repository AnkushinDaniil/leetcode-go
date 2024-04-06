package kokoEatingBananas

func minEatingSpeed(piles []int, h int) int {
	minimum := 1
	maximum := piles[0]

	if len(piles) > 1 {
		for _, bananas := range piles[1:] {
			if bananas > maximum {
				maximum = bananas
			}
		}
	}

	for minimum <= maximum {
		mid := (minimum + maximum) / 2
		hours := 0
		for _, bananas := range piles {
			hours += (bananas + mid - 1) / mid
		}
		if hours > h {
			minimum = mid + 1
		} else {
			maximum = mid - 1
		}
	}

	return minimum
}
