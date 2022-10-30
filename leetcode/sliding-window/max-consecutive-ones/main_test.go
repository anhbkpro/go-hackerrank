package max_consecutive_ones

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxConsecutiveOnes(t *testing.T) {
	assert.Equal(t, 3, FindMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	assert.Equal(t, 2, FindMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}
