package sort_array_by_increasing_frequency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	assert.Equal(t, []int{1, 3, 3, 2, 2}, FrequencySort([]int{2, 3, 1, 3, 2}))
	assert.Equal(t, []int{3, 1, 1, 2, 2, 2}, FrequencySort([]int{1, 1, 2, 2, 2, 3}))
	assert.Equal(t, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}, FrequencySort([]int{-1, 1, -6, 4, 5, -6, 1, 4, 1}))
}
