package coinChange

func coinChange(coins []int, amount int) int {
	n := amount + 1
	memo := make([]int, n)
	for i := 1; i < n; i++ {
		memo[i] = n
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < n; j++ {
			memo[j] = min(memo[j], memo[j-coins[i]]+1)
		}
	}
	if memo[amount] == amount+1 {
		return -1
	}

	return memo[amount]
}
