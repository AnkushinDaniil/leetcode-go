package coinChange

var mem [10001]int

func coinChange(coins []int, amount int) int {
	n := amount + 1

	for i := range mem[1:n] {
		mem[i+1] = amount + 1
	}

	for _, coin := range coins {
		for j := coin; j < n; j++ {
			mem[j] = min(mem[j], mem[j-coin]+1)
		}
	}
	if mem[amount] == n {
		return -1
	}
	return mem[amount]
}
