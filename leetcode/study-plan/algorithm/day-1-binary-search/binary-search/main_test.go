package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, Search([]int{-1, 0, 3, 5, 9, 12}, 9))
}
