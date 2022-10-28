package k_radius_subarray_averages

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAverages(t *testing.T) {
	assert.Equal(t, []int{100_000}, GetAverages([]int{100_000}, 0))
	assert.Equal(t, []int{-1}, GetAverages([]int{8}, 100_000))
	assert.Equal(t, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}, GetAverages([]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3))
}
