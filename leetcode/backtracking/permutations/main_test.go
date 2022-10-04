package permutations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	got := Permute([]int{1, 2, 3})
	assert.Equal(t, [][]int{{3, 2, 1}, {2, 3, 1}, {1, 3, 2}, {3, 1, 2}, {2, 1, 3}, {1, 2, 3}}, got)
}
