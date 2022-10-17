package best_time_to_buy_and_sell_stock_with_cooldown

import (
	"fmt"
	"math"
)

type Transaction struct {
	index  int
	canBuy bool
}

func MaxProfit(prices []int) int {
	// State: Buying or Selling?
	// If Buy => i + 1
	// If Sell => i + 2
	dp := make(map[Transaction]int) // key=(index, canBuy) value=max_profit
	return dfs(&dp, prices, 0, true)
}

func dfs(dp *map[Transaction]int, prices []int, i int, canBuy bool) int {
	if i >= len(prices) {
		return 0
	}

	t := Transaction{index: i, canBuy: canBuy}
	if _, ok := (*dp)[t]; ok {
		return (*dp)[t]
	}

	coolDown := dfs(dp, prices, i+1, canBuy)
	if canBuy {
		buy := dfs(dp, prices, i+1, !canBuy) - prices[i]
		max := int(math.Max(float64(buy), float64(coolDown)))
		fmt.Printf("[ BUY or COOLDOWN] dp(%d, %v)=%d\n", i, canBuy, max)
		(*dp)[Transaction{index: i, canBuy: canBuy}] = max
	} else {
		sell := dfs(dp, prices, i+2, !canBuy) + prices[i]
		max := int(math.Max(float64(sell), float64(coolDown)))
		fmt.Printf("[SELL or COOLDOWN] dp(%d, %v)=%d\n", i, canBuy, max)
		(*dp)[Transaction{index: i, canBuy: canBuy}] = max
	}
	return (*dp)[Transaction{index: i, canBuy: canBuy}]
}
