package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 2, ClimbStairs(2))
	assert.Equal(t, 3, ClimbStairs(3))
	assert.Equal(t, 5, ClimbStairs(4))
	assert.Equal(t, 5, ClimbStairs(4))
}
