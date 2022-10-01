package min_cost_climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	got := MinCostClimbingStairs([]int32{10, 15, 20})
	assert.Equal(t, int32(15), got)
}
