package find_the_most_competitive_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMostCompetitive(t *testing.T) {
	assert.Equal(t, []int{10, 23, 61, 62, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, MostCompetitive([]int{84, 10, 71, 23, 66, 61, 62, 64, 34, 41, 80, 25, 91, 43, 4, 75, 65, 13, 37, 41, 46, 90, 55, 8, 85, 61, 95, 71}, 24))
	assert.Equal(t, []int{8, 80, 2}, MostCompetitive([]int{71, 18, 52, 29, 55, 73, 24, 42, 66, 8, 80, 2}, 3))
	assert.Equal(t, []int{2, 6}, MostCompetitive([]int{3, 5, 2, 6}, 2))
	assert.Equal(t, []int{2, 3, 3, 4}, MostCompetitive([]int{2, 4, 3, 3, 5, 4, 9, 6}, 4))
}
