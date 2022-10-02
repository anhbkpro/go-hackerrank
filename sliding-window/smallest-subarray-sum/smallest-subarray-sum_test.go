package smallest_subarray_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSizeSubArraySum(t *testing.T) {
	input := []int32{2, 1, 5, 2, 3, 2}
	got := MinSizeSubArraySum(7, input)
	assert.Equal(t, int32(2), got)
}
