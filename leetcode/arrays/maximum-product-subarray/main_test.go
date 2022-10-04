package maximum_product_subarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	got := MaxProduct([]int{2, 3, -2, 4})
	assert.Equal(t, 6, got)

	got2 := MaxProduct([]int{-2})
	assert.Equal(t, -2, got2)

	got3 := MaxProduct([]int{-3, -1, -1})
	assert.Equal(t, 3, got3)

	got4 := MaxProduct([]int{0, 2})
	assert.Equal(t, 2, got4)

	got5 := MaxProduct([]int{-2, 3, -4})
	assert.Equal(t, 24, got5)
}
