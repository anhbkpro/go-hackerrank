package consecutive_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxPower(t *testing.T) {
	assert.Equal(t, 3, MaxPower("ccddd"))
	assert.Equal(t, 2, MaxPower("cc"))
	assert.Equal(t, 1, MaxPower("c"))
	assert.Equal(t, 2, MaxPower("leetcode"))
	assert.Equal(t, 5, MaxPower("abbcccddddeeeeedcba"))
}
