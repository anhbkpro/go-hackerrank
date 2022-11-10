package top_k_frequent_elements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	assert.Equal(t, []int{1, 2}, TopKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	assert.Equal(t, []int{1}, TopKFrequent([]int{1}, 1))
}
