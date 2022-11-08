package longest_string_chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestStrChain(t *testing.T) {
	assert.Equal(t, 4, LongestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	assert.Equal(t, 5, LongestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
}
