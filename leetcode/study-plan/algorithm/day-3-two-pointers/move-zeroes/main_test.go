package move_zeroes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	arr := []int{0, 1, 0, 3, 12}
	MoveZeroes(arr)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, arr)
}
