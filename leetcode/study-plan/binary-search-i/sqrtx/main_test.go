package sqrtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMySqrt(t *testing.T) {
	assert.Equal(t, 1, MySqrt(1))
	assert.Equal(t, 2, MySqrt(8))
	assert.Equal(t, 3, MySqrt(11))
}
