package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	got := MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(t, 49, got)
}
