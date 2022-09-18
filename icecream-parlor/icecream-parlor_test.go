package icecream_parlor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIcecreamParlor(t *testing.T) {
	input := []int32{1, 4, 5, 3, 2}
	got := IcecreamParlor(4, input)
	assert.Equal(t, []int32{1, 4}, got)
}
