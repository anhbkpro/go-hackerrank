package _32_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind132pattern(t *testing.T) {
	assert.Equal(t, true, Find132pattern([]int{9, 11, 8, 9, 10, 7, 9}))
	assert.Equal(t, true, Find132pattern([]int{3, 1, 4, 2}))
	assert.Equal(t, false, Find132pattern([]int{1, 2, 3, 4}))
	assert.Equal(t, true, Find132pattern([]int{-1, 3, 2, 0}))
}
