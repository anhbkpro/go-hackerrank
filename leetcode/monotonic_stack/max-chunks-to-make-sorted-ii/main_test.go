package max_chunks_to_make_sorted_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	assert.Equal(t, 1, MaxChunksToSorted([]int{5, 4, 3, 2, 1}))
	assert.Equal(t, 4, MaxChunksToSorted([]int{2, 1, 3, 4, 4}))
}
