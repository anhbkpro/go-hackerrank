package sum_of_subarray_ranges

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubArrayRanges(t *testing.T) {
	assert.Equal(t, int64(4), SubArrayRanges([]int{1, 2, 3}))
}
