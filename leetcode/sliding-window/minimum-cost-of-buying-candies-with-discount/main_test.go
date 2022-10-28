package minimum_cost_of_buying_candies_with_discount

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumCost(t *testing.T) {
	assert.Equal(t, 5, MinimumCost([]int{1, 2, 3}))
	assert.Equal(t, 23, MinimumCost([]int{6, 5, 7, 9, 2, 2}))
	assert.Equal(t, 10, MinimumCost([]int{5, 5}))
}
