package best_time_to_buy_and_sell_stock_with_cooldown

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 3, MaxProfit([]int{1, 2, 3, 0, 2}))
	assert.Equal(t, 0, MaxProfit([]int{1}))
}
