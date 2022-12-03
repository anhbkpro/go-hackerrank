package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) int {
	l, r := 0, 1 //l = buy, right = sell
	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = max(maxP, profit)
		} else {
			l = r
		}
		r++
	}
	return maxP
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
