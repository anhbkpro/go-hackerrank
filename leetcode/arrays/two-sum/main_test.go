package two_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, []int{0, 1}, got)
}
