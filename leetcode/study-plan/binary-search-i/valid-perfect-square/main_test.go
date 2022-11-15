package valid_perfect_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	assert.Equal(t, true, IsPerfectSquare(1))
	assert.Equal(t, true, IsPerfectSquare(16))
	assert.Equal(t, false, IsPerfectSquare(15))
}
