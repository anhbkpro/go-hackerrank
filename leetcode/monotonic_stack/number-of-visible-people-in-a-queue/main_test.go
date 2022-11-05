package number_of_visible_people_in_a_queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanSeePersonsCount(t *testing.T) {
	assert.Equal(t, []int{3, 1, 2, 1, 1, 0}, CanSeePersonsCount([]int{10, 6, 8, 5, 11, 9}))
	assert.Equal(t, []int{4, 1, 1, 1, 0}, CanSeePersonsCount([]int{5, 1, 2, 3, 10}))
}
