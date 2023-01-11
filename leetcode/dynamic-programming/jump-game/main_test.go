package jump_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanJump(t *testing.T) {
	assert.Equal(t, false, CanJumpGreedy([]int{1, 0, 1, 0}))
	assert.Equal(t, true, CanJumpGreedy([]int{0}))
	assert.Equal(t, false, CanJumpGreedy([]int{0, 2, 3}))
	assert.Equal(t, true, CanJumpGreedy([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, CanJumpGreedy([]int{3, 2, 1, 0, 4}))
}
