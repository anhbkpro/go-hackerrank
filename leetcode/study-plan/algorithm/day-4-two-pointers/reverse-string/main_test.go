package reverse_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseString(arr)
	assert.Equal(t, []byte{'o', 'l', 'l', 'e', 'h'}, arr)
}
