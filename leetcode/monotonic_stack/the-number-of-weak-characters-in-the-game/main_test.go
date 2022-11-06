package the_number_of_weak_characters_in_the_game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfWeakCharacters(t *testing.T) {
	assert.Equal(t, 6, NumberOfWeakCharacters([][]int{{7, 7}, {1, 2}, {9, 7}, {7, 3}, {3, 10}, {9, 8}, {8, 10}, {4, 3}, {1, 5}, {1, 5}}))
	assert.Equal(t, 0, NumberOfWeakCharacters([][]int{{5, 5}, {6, 3}, {3, 6}}))
	assert.Equal(t, 1, NumberOfWeakCharacters([][]int{{2, 2}, {3, 3}}))
	assert.Equal(t, 1, NumberOfWeakCharacters([][]int{{1, 5}, {10, 4}, {4, 3}}))
}
