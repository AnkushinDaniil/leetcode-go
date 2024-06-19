package gasStation

func canCompleteCircuit(gas []int, cost []int) int {
	sum := 0
	curSum := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		dif := gas[i] - cost[i]
		sum += dif
		curSum += dif
		if curSum < 0 {
			curSum = 0
			start = i + 1
		}
	}

	if sum < 0 {
		return -1
	} else {
		return start
	}
}

func canCompleteCircuitLCT(gas []int, cost []int) int {
	var totalTank, currTank, start int

	for i := range gas {
		totalTank += gas[i] - cost[i]
		currTank += gas[i] - cost[i]

		if currTank < 0 {
			// means we cannot start from this station, so start from next
			start = i + 1
			currTank = 0
		}
	}

	if totalTank >= 0 {
		return start
	}

	return -1
}

func canCompleteCircuitLCM(gas []int, cost []int) int {
	if len(gas) == 1 {
		if gas[0] >= cost[0] {
			return 0
		}
		return -1
	}
	// diff, c, index := math.MinInt64, -1, -1
	for i := 0; i < len(gas); i++ {
		if gas[i] > cost[i] {
			if getTank(i, gas, cost) >= 0 {
				return i
			}
		}
	}
	// if index == -1{
	//     return -1
	// }
	return -1
}

func getTank(index int, gas, cost []int) int {
	tank := 0
	curIndex := index
	for i := 0; i < len(gas); i++ {
		if curIndex == index {
			tank = tank + gas[curIndex]
			curIndex++
			continue
		}
		if curIndex == len(gas) || curIndex == 0 {
			curIndex = 0
			if tank-cost[len(cost)-1] <= 0 {
				return -1
			}
			tank = tank - cost[len(cost)-1] + gas[curIndex]
			curIndex++
			continue
		}
		if tank-cost[curIndex-1] <= 0 {
			return -1
		}
		tank = tank - cost[curIndex-1] + gas[curIndex]
		curIndex++
	}
	if tank > 0 {
		if index == 0 {
			tank = tank - cost[len(cost)-1]
		} else {
			tank = tank - cost[index-1]
		}
	}
	if tank >= 0 {
		return index
	}
	return -1
}
