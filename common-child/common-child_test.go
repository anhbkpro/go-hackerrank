package common_child

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonChild(t *testing.T) {
	got1 := CommonChild("HARRY", "SALLY")
	assert.Equal(t, int32(2), got1)

	got2 := CommonChild("AA", "BB")
	assert.Equal(t, int32(0), got2)

	got3 := CommonChild("SHINCHAN", "NOHARAAA")
	assert.Equal(t, int32(3), got3)
}
