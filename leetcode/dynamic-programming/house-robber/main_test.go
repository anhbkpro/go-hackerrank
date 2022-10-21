package house_robber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRob(t *testing.T) {
	assert.Equal(t, 4, Rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 12, Rob([]int{2, 7, 9, 3, 1}))
}
