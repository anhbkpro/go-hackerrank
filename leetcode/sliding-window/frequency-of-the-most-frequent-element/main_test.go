package frequency_of_the_most_frequent_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	assert.Equal(t, 3, MaxFrequency([]int{1, 2, 4}, 5))
	assert.Equal(t, 2, MaxFrequency([]int{1, 4, 8, 13}, 5))
	assert.Equal(t, 1, MaxFrequency([]int{3, 9, 6}, 2))
}
