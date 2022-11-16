package find_the_distance_value_between_two_arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTheDistanceValue(t *testing.T) {
	assert.Equal(t, 2, FindTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2))
}
