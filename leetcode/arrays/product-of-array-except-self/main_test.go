package product_of_array_except_self

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	got := ProductExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(t, []int{24, 12, 8, 6}, got)
}
