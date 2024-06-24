package reconstructItinerary

import (
	"slices"
	"sort"
	"strings"
)

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

var (
	start        = "JFK"
	stackArr     = [301]*string{&start}
	itineraryArr = [301]string{}
)

func findItineraryPointers(tickets [][]string) []string {
	nodes := make(map[string][]*string)
	for i := 0; i < len(tickets); i++ {
		nodes[tickets[i][0]] = append(nodes[tickets[i][0]], &tickets[i][1])
	}

	for key := range nodes {
		slices.SortFunc(nodes[key], func(a, b *string) int {
			return strings.Compare(*a, *b)
		})
	}

	stack := stackArr[:1]
	itinerary := itineraryArr[:]
	i := 0

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		if len(nodes[*cur]) > 0 {
			next := nodes[*cur][0]
			nodes[*cur] = nodes[*cur][1:]
			stack = append(stack, next)
		} else {
			itinerary[len(tickets)-i] = *cur
			i++
			stack = stack[:len(stack)-1]
		}
	}

	return itinerary[:i]
}

type keys []byte

func (a keys) Len() int { return len(a) / 3 }
func (a keys) Swap(i, j int) {
	for k := 0; k < 3; k++ {
		a[i*3+k], a[j*3+k] = a[j*3+k], a[i*3+k]
	}
}

func (a keys) Less(i, j int) bool {
	for k := 0; k < 3; k++ {
		if a[i*3+k] == a[j*3+k] {
			continue
		} else {
			return a[i*3+k] < a[j*3+k]
		}
	}
	return false
}

func findItineraryString(tickets [][]string) []string {
	nodes := make(map[string]keys, 4)
	for i := 0; i < len(tickets); i++ {
		nodes[tickets[i][0]] = append(nodes[tickets[i][0]], []byte(tickets[i][1])...)
	}

	for key := range nodes {
		sort.Sort(nodes[key])
	}

	stack := make([]byte, 0, (len(tickets)+1)*3)
	stack = append(stack, "JFK"...)
	itinerary := make([]string, 301)
	i := 0

	for len(stack) > 0 {
		cur := string(stack[len(stack)-3:])
		if len(nodes[cur]) > 0 {
			next := nodes[cur][:3]
			nodes[cur] = nodes[cur][3:]
			stack = append(stack, next...)
		} else {
			itinerary[len(tickets)-i] = cur
			i++
			stack = stack[:len(stack)-3]
		}
	}

	return itinerary[:i]
}

func findItineraryLCT(tickets [][]string) []string {
	adjList := map[string][]string{}

	for _, ticket := range tickets {
		adjList[ticket[0]] = append(adjList[ticket[0]], ticket[1])
	}

	for _, adj := range adjList {
		sort.Strings(adj)
	}

	s := solution{
		adjList: adjList,
		result:  []string{},
	}
	s.dfs("JFK")

	slices.Reverse(s.result)
	return s.result
}

type solution struct {
	adjList map[string][]string
	result  []string
}

func (s *solution) dfs(city string) {
	for len(s.adjList[city]) > 0 {
		next := s.adjList[city][0]
		s.adjList[city] = s.adjList[city][1:]
		s.dfs(next)
	}
	s.result = append(s.result, city)
}
