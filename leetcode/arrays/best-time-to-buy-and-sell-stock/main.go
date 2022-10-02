package best_time_to_buy_and_sell_stock

import "math"

func MaxProfit(prices []int) int {
	l, r := 0, 1 //l = buy, right = sell
	maxP := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxP = int(math.Max(float64(maxP), float64(profit)))
		} else {
			l = r
		}
		r++
	}
	return maxP
}
