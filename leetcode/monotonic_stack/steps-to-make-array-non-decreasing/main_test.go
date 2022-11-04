package steps_to_make_array_non_decreasing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalSteps(t *testing.T) {
	assert.Equal(t, 3, TotalSteps([]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}))
	assert.Equal(t, 0, TotalSteps([]int{4, 5, 7, 7, 13}))
}
