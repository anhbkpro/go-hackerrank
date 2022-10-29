package maximum_erasure_value

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaximumUniqueSubarray(t *testing.T) {
	assert.Equal(t, 17, MaximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
	assert.Equal(t, 8, MaximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
}
