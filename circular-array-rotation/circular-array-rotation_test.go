package circular_array_rotation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircularArrayRotation(t *testing.T) {
	got := CircularArrayRotation([]int32{1, 2, 3}, 2, []int32{0, 1, 2})
	// {1, 2, 3} ---k=1---> {3, 1, 2} ---k=2---> {2, 3, 1}
	assert.Equal(t, []int32{2, 3, 1}, got)
}
