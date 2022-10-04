package subarray_product_less_than_k

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	got := NumSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	assert.Equal(t, 8, got)
}
