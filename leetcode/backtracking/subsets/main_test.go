package subsets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	got := Subsets([]int{1, 2, 3})
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}}, got)

	got2 := Subsets([]int{3, 2, 4, 1})
	assert.Equal(t, [][]int{
		{3, 2, 4, 1},
		{3, 2, 4},
		{3, 2, 1},
		{3, 2},
		{3, 4, 1},
		{3, 4},
		{3, 1},
		{3},
		{2, 4, 1},
		{2, 4},
		{2, 1},
		{2},
		{4, 1},
		{4},
		{1},
		{}}, got2)
}
