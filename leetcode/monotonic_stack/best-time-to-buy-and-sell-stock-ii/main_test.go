package best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 7, MaxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, MaxProfit([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, MaxProfit([]int{7, 6, 4, 3, 1}))
}
