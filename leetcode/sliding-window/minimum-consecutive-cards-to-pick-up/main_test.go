package minimum_consecutive_cards_to_pick_up

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinimumCardPickup(t *testing.T) {
	assert.Equal(t, 4, MinimumCardPickup([]int{3, 4, 2, 3, 4, 7}))
	assert.Equal(t, -1, MinimumCardPickup([]int{1, 0, 5, 3}))
}
