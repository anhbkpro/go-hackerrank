package minimum_time_to_make_rope_colorful

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinCost(t *testing.T) {
	got := MinCost("abaac", []int{1, 2, 3, 4, 5})
	assert.Equal(t, 3, got)

	got2 := MinCost("aabaa", []int{1, 2, 3, 4, 1})
	assert.Equal(t, 2, got2)

	got3 := MinCost("bbbaaa", []int{4, 9, 3, 8, 8, 9})
	assert.Equal(t, 23, got3)
}
