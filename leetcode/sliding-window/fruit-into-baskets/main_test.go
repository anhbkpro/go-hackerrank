package fruit_into_baskets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTotalFruit(t *testing.T) {
	assert.Equal(t, 3, TotalFruit([]int{1, 2, 1}))
	assert.Equal(t, 3, TotalFruit([]int{0, 1, 2, 2}))
	assert.Equal(t, 4, TotalFruit([]int{1, 2, 3, 2, 2}))
}
