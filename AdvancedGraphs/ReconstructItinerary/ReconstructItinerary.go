package reconstructItinerary

import "slices"

func findItinerary(tickets [][]string) []string {
	nodes := make(map[string][]string, 4)
	for i := 0; i < len(tickets); i++ {
		if _, ok := nodes[tickets[i][0]]; !ok {
			nodes[tickets[i][0]] = make([]string, 0, 4)
		}
		nodes[tickets[i][0]] = append(nodes[tickets[i][0]], tickets[i][1])
	}

	for key := range nodes {
		slices.Sort(nodes[key])
	}

	stack := make([]string, 0, len(tickets)+1)
	stack = append(stack, "JFK")
	itinerary := make([]string, len(tickets)+1)
	i := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(nodes[cur]) > 0 {
			next := nodes[cur][0]
			nodes[cur] = nodes[cur][1:]
			stack = append(stack, next)
		} else {
			itinerary[len(itinerary)-1-i] = cur
			i++
			stack = stack[:len(stack)-1]
		}
	}

	return itinerary
}
