package minimum_size_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	assert.Equal(t, 2, MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, MinSubArrayLen(4, []int{1, 4, 4}))
	assert.Equal(t, 0, MinSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
