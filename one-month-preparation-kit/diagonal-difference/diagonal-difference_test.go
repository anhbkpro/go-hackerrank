package diagonal_difference

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	input := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	got := DiagonalDifference(input)
	assert.Equal(t, int32(15), got)
}
