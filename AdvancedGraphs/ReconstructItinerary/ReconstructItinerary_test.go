package reconstructItinerary

import (
	"fmt"
	"testing"

	"github.com/test-go/testify/assert"
)

var testTable = []struct {
	tickets [][]string
	output  []string
}{
	{
		tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
		output:  []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
	},
	{
		tickets: [][]string{
			{"JFK", "SFO"},
			{"JFK", "ATL"},
			{"SFO", "ATL"},
			{"ATL", "JFK"},
			{"ATL", "SFO"},
		},
		output: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
	},
	{
		tickets: [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}},
		output:  []string{"JFK", "NRT", "JFK", "KUL"},
	},
	{
		tickets: [][]string{
			{"BIM", "VIE"},
			{"ANU", "EZE"},
			{"PER", "CNS"},
			{"TCB", "BIM"},
			{"MEL", "PER"},
			{"VIE", "ANU"},
			{"ANU", "SYD"},
			{"SYD", "PER"},
			{"DRW", "TIA"},
			{"PER", "BAK"},
			{"EZE", "BNE"},
			{"ANU", "ADL"},
			{"BIM", "MEL"},
			{"EZE", "ASD"},
			{"SYD", "JFK"},
			{"ADL", "BIM"},
			{"AUA", "EZE"},
			{"CNS", "TBI"},
			{"DRW", "ANU"},
			{"OOL", "VIE"},
			{"DRW", "TCB"},
			{"OOL", "AXA"},
			{"TIA", "BIM"},
			{"TIA", "ANU"},
			{"ANU", "MEL"},
			{"TIA", "MHH"},
			{"SYD", "AXA"},
			{"AXA", "ADL"},
			{"SYD", "ANU"},
			{"CBR", "ADL"},
			{"CBR", "TIA"},
			{"ADL", "BIM"},
			{"TBI", "DRW"},
			{"MHH", "ADL"},
			{"MEL", "HBA"},
			{"CNS", "JFK"},
			{"VIE", "CBR"},
			{"ANU", "VIE"},
			{"BNE", "HBA"},
			{"BIM", "TCB"},
			{"JFK", "BAK"},
			{"EZE", "SYD"},
			{"AUA", "BIM"},
			{"TBI", "TIA"},
			{"MEL", "JFK"},
			{"TIA", "BAK"},
			{"BNE", "VIE"},
			{"HBA", "PER"},
			{"ANU", "AUA"},
			{"EZE", "BNE"},
			{"INN", "AUA"},
			{"TIA", "ANU"},
			{"BAK", "EZE"},
			{"INN", "JFK"},
			{"BAK", "MEL"},
			{"AUA", "TCB"},
			{"PER", "BAK"},
			{"SYD", "DRW"},
			{"LST", "MEL"},
			{"INN", "BNE"},
			{"LST", "JFK"},
			{"AXA", "MEL"},
			{"BAK", "CNS"},
			{"BNE", "HBA"},
			{"VIE", "TCB"},
			{"MHH", "ASD"},
			{"OOL", "PER"},
			{"PER", "EZE"},
			{"MEL", "AXA"},
			{"BNE", "SYD"},
			{"MEL", "ANU"},
			{"LST", "DRW"},
			{"EZE", "TBI"},
			{"ASD", "AXA"},
			{"MHH", "LST"},
			{"TIA", "EZE"},
			{"DRW", "HBA"},
			{"SYD", "BGI"},
			{"AUA", "MHH"},
			{"BIM", "INN"},
			{"AXA", "VIE"},
			{"EZE", "CNS"},
			{"JFK", "BNE"},
			{"TCB", "CNS"},
			{"ADL", "CBR"},
			{"BIM", "INN"},
			{"TBI", "TCB"},
			{"BIM", "MEL"},
			{"VIE", "BNE"},
			{"JFK", "TIA"},
			{"ADL", "AUA"},
			{"ANU", "OOL"},
			{"SYD", "AUA"},
			{"JFK", "AUA"},
			{"HBA", "BIM"},
			{"SYD", "ADL"},
			{"TCB", "DRW"},
			{"JFK", "BNE"},
			{"MHH", "SYD"},
			{"CNS", "EZE"},
			{"AXA", "TIA"},
			{"INN", "SYD"},
			{"CBR", "INN"},
			{"ADL", "TIA"},
			{"SYD", "LST"},
			{"BAK", "TIA"},
			{"DRW", "INN"},
			{"CBR", "BAK"},
			{"ASD", "SYD"},
			{"EZE", "ADL"},
			{"JFK", "LST"},
			{"BNE", "OOL"},
			{"SYD", "BIM"},
			{"ADL", "EZE"},
			{"BNE", "VIE"},
			{"BAK", "PER"},
			{"BNE", "VIE"},
			{"EZE", "ASD"},
			{"BAK", "EZE"},
			{"EZE", "JFK"},
			{"LST", "BNE"},
			{"VIE", "SYD"},
			{"BNE", "VIE"},
			{"BAK", "AUA"},
			{"ASD", "LST"},
			{"VIE", "SYD"},
			{"OOL", "ADL"},
			{"EZE", "SYD"},
			{"AUA", "EZE"},
			{"ASD", "AXA"},
			{"AXA", "OOL"},
			{"TCB", "EZE"},
			{"TBI", "JFK"},
			{"HBA", "SYD"},
			{"AXA", "CBR"},
			{"INN", "BNE"},
			{"BIM", "TBI"},
			{"ANU", "OOL"},
			{"DRW", "MHH"},
			{"BNE", "ANU"},
			{"TCB", "OOL"},
			{"TIA", "CBR"},
			{"OOL", "MHH"},
			{"VIE", "ASD"},
			{"VIE", "ADL"},
			{"HBA", "DRW"},
			{"MEL", "EZE"},
			{"PER", "AXA"},
			{"ADL", "SYD"},
			{"CNS", "BNE"},
			{"JFK", "TBI"},
			{"SYD", "DRW"},
			{"MEL", "ANU"},
			{"JFK", "BAK"},
			{"EZE", "BAK"},
			{"AUA", "EZE"},
			{"PER", "INN"},
			{"EZE", "MEL"},
			{"ADL", "PER"},
		},
		output: []string{
			"JFK",
			"AUA",
			"BIM",
			"INN",
			"AUA",
			"EZE",
			"ADL",
			"AUA",
			"EZE",
			"ASD",
			"AXA",
			"ADL",
			"BIM",
			"INN",
			"BNE",
			"ANU",
			"ADL",
			"BIM",
			"MEL",
			"ANU",
			"AUA",
			"EZE",
			"ASD",
			"AXA",
			"CBR",
			"ADL",
			"CBR",
			"BAK",
			"AUA",
			"MHH",
			"ADL",
			"EZE",
			"BAK",
			"CNS",
			"BNE",
			"HBA",
			"BIM",
			"MEL",
			"ANU",
			"EZE",
			"BNE",
			"HBA",
			"DRW",
			"ANU",
			"MEL",
			"AXA",
			"MEL",
			"EZE",
			"BNE",
			"OOL",
			"ADL",
			"PER",
			"AXA",
			"OOL",
			"AXA",
			"TIA",
			"ANU",
			"OOL",
			"MHH",
			"ASD",
			"LST",
			"BNE",
			"SYD",
			"ADL",
			"SYD",
			"ANU",
			"OOL",
			"PER",
			"BAK",
			"EZE",
			"CNS",
			"EZE",
			"JFK",
			"BAK",
			"EZE",
			"MEL",
			"HBA",
			"PER",
			"BAK",
			"MEL",
			"JFK",
			"BAK",
			"PER",
			"CNS",
			"JFK",
			"BNE",
			"VIE",
			"ADL",
			"TIA",
			"ANU",
			"SYD",
			"AUA",
			"TCB",
			"BIM",
			"TBI",
			"DRW",
			"HBA",
			"SYD",
			"AXA",
			"VIE",
			"ANU",
			"VIE",
			"ASD",
			"SYD",
			"BIM",
			"TCB",
			"CNS",
			"TBI",
			"JFK",
			"BNE",
			"VIE",
			"BNE",
			"VIE",
			"CBR",
			"INN",
			"BNE",
			"VIE",
			"SYD",
			"DRW",
			"INN",
			"JFK",
			"LST",
			"DRW",
			"MHH",
			"LST",
			"JFK",
			"TBI",
			"TCB",
			"DRW",
			"TCB",
			"EZE",
			"SYD",
			"DRW",
			"TIA",
			"BAK",
			"TIA",
			"BIM",
			"VIE",
			"TCB",
			"OOL",
			"VIE",
			"SYD",
			"JFK",
			"TIA",
			"CBR",
			"TIA",
			"EZE",
			"SYD",
			"LST",
			"MEL",
			"PER",
			"EZE",
			"TBI",
			"TIA",
			"MHH",
			"SYD",
			"PER",
			"INN",
			"SYD",
			"BGI",
		},
	},
	{
		tickets: [][]string{
			{"JFK", "SFO"},
			{"JFK", "ATL"},
			{"SFO", "JFK"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
			{"ATL", "AAA"},
			{"AAA", "BBB"},
			{"BBB", "ATL"},
		},
		output: []string{
			"JFK",
			"SFO",
			"JFK",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
			"AAA",
			"BBB",
			"ATL",
		},
	},
}

func Test(t *testing.T) {
	for i := 0; i < len(testTable); i++ {
		t.Run(fmt.Sprintf("Test %v", testTable[i].tickets), func(t *testing.T) {
			assert.Equal(
				t,
				testTable[i].output,
				findItineraryPointers(testTable[i].tickets),
			)
		})
	}
}

func Benchmark_findItinerary(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findItinerary(testTable[4].tickets)
	}
}

func Benchmark_findItineraryPointers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findItineraryPointers(testTable[4].tickets)
	}
}

func Benchmark_findItineraryString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findItineraryString(testTable[4].tickets)
	}
}

func Benchmark_findItineraryLCT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findItineraryLCT(testTable[4].tickets)
	}
}

// Benchmark_findItinerary-10                101001             11190 ns/op       5824 B/op          16 allocs/op
// Benchmark_findItineraryPointers-10        125648              9433 ns/op       1544 B/op          21 allocs/op
// Benchmark_findItineraryString-10           82164             14435 ns/op       6528 B/op         187 allocs/op
// Benchmark_findItineraryLCT-10             104736             10996 ns/op       7952 B/op          31 allocs/op
