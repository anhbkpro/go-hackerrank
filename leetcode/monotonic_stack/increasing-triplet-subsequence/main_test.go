package increasing_triplet_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	assert.Equal(t, true, IncreasingTriplet([]int{20, 100, 10, 12, 5, 13}))
	assert.Equal(t, true, IncreasingTriplet([]int{1, 5, 0, 4, 1, 3}))
	assert.Equal(t, false, IncreasingTriplet([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, true, IncreasingTriplet([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, true, IncreasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}
