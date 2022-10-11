package combination_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, CombinationSum([]int{2, 3, 6, 7}, 7))
}
