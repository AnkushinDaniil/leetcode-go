package carFleet

// import "sort"

// type Car struct {
// 	position int
// 	speed    int
// }

// type Cars []Car

// func (c Cars) Len() int           { return len(c) }
// func (c Cars) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
// func (c Cars) Less(i, j int) bool { return c[i].position < c[j].position }

// func carFleet(target int, position []int, speed []int) int {
// 	cars := make(Cars, len(position))
// 	for i := range position {
// 		cars[i] = Car{
// 			position: position[i],
// 			speed:    speed[i],
// 		}
// 	}
// 	sort.Sort(cars)

// 	stack := make(Cars, 1)
// 	stack[0] = cars[cars.Len()-1]

// 	for i := cars.Len() - 2; i >= 0; i-- {
// 		topStackCar := stack[len(stack)-1]
// 		car := cars[i]

// 		if topStackCar.speed < car.speed {
// 			if (topStackCar.position-car.position)*topStackCar.speed <=
// 				(target-topStackCar.position)*(car.speed-topStackCar.speed) {
// 				continue
// 			}
// 		}

// 		stack = append(stack, car)
// 	}

// 	return stack.Len()
// }

func carFleet(target int, position []int, speed []int) int {
	positionSpeed := make([]int, target+1)

	for i := 0; i < len(position); i++ {
		positionSpeed[position[i]] = speed[i]
	}

	var lastCatPosition, lastCarSpeed int
	i := target

	for ; i >= 0; i-- {
		if positionSpeed[i] == 0 {
			continue
		} else {
			lastCatPosition = i
			lastCarSpeed = positionSpeed[i]
			break
		}
	}

	ans := 1

	for ; i >= 0; i-- {
		if positionSpeed[i] == 0 {
			continue
		}

		if (target-lastCatPosition)*positionSpeed[i] < (target-i)*lastCarSpeed {
			lastCatPosition = i
			lastCarSpeed = positionSpeed[i]
			ans++
		}
	}

	return ans
}
