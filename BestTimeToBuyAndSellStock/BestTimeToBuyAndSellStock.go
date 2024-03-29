package bestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	mp := 0

	for l, r := 0, 1; r < len(prices); r++ {
		if prices[l] > prices[r] {
			l = r
		} else if profit := prices[r] - prices[l]; profit > mp {
			mp = profit
		}
	}

	return mp
}
