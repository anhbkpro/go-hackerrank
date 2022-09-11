package service_lane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServiceLane(t *testing.T) {
	width := []int32{2, 3, 1, 2, 3, 2, 3, 3}
	cases := [][]int32{
		{0, 3},
		{4, 6},
		{6, 7},
		{3, 5},
		{0, 7},
	}
	got := ServiceLane(8, width, cases)
	assert.Equal(t, []int32{1, 2, 3, 2, 1}, got)

	width2 := []int32{1, 2, 2, 2, 1}
	cases2 := [][]int32{
		{2, 3},
		{1, 4},
		{2, 4},
		{2, 4},
		{2, 3},
	}
	got2 := ServiceLane(5, width2, cases2)
	assert.Equal(t, []int32{2, 1, 1, 1, 2}, got2)
}
