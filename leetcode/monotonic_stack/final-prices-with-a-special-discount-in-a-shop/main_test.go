package final_prices_with_a_special_discount_in_a_shop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFinalPrices(t *testing.T) {
	assert.Equal(t, []int{4, 2, 4, 2, 3}, FinalPrices([]int{8, 4, 6, 2, 3}))
	assert.Equal(t, []int{9, 0, 1, 6}, FinalPrices([]int{10, 1, 1, 6}))
}
