package peak_index_in_a_mountain_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeakIndexInMountainArray(t *testing.T) {
	assert.Equal(t, 1, PeakIndexInMountainArray([]int{0, 2, 1, 0}))
}
