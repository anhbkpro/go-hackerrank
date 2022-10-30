package max_consecutive_ones_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	assert.Equal(t, 6, LongestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	assert.Equal(t, 10, LongestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
}
