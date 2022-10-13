package longest_increasing_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	assert.Equal(t, 4, LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
