package find_all_lonely_numbers_in_the_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLonely(t *testing.T) {
	assert.Equal(t, []int{10, 8}, FindLonely([]int{10, 6, 5, 8}))
	assert.Equal(t, []int{1, 5}, FindLonely([]int{1, 3, 5, 3}))
}
