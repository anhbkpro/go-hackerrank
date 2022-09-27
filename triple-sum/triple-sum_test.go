package triple_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTriplets(t *testing.T) {
	got := Triplets([]int32{1, 4, 5}, []int32{2, 3, 3}, []int32{1, 2, 3})
	assert.Equal(t, int64(5), got)
}
