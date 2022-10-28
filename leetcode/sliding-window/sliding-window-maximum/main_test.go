package sliding_window_maximum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal(t, []int{1}, MaxSlidingWindow([]int{1}, 1))
}
