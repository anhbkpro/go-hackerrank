package one_month_preparation_kit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLonelyinteger(t *testing.T) {
	input := []int32{1, 2, 3, 4, 3, 2, 1}
	got := Lonelyinteger(input)
	assert.Equal(t, int32(4), got)
}
