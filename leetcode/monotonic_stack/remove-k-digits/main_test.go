package remove_k_digits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveKdigits(t *testing.T) {
	assert.Equal(t, "1219", RemoveKdigits("1432219", 3))
	assert.Equal(t, "200", RemoveKdigits("10200", 1))
	assert.Equal(t, "0", RemoveKdigits("10", 2))
}
