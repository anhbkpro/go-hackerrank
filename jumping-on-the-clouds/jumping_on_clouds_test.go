package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJumpingOnClouds(t *testing.T) {
	input := []int32{0, 0, 1, 0, 0, 1, 0}
	got := JumpingOnClouds(input)
	assert.Equal(t, int32(4), got)

	input2 := []int32{0, 0, 0, 1, 0, 0}
	got2 := JumpingOnClouds(input2)
	assert.Equal(t, int32(3), got2)

	input3 := []int32{0, 0}
	got3 := JumpingOnClouds(input3)
	assert.Equal(t, int32(1), got3)

	input4 := []int32{1, 0}
	got4 := JumpingOnClouds(input4)
	assert.Equal(t, int32(1), got4)
}
