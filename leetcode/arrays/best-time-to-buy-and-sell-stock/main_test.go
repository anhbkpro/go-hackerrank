package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	got := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 5, got)
}
