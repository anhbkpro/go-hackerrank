package max_chunks_to_make_sorted

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {
	// The first chunk can be found as the smallest k for which A[:k+1] == [0, 1, 2, ...k]; then we repeat this process.
	assert.Equal(t, 2, MaxChunksToSorted([]int{1, 2, 0, 3}))
	assert.Equal(t, 4, MaxChunksToSorted([]int{1, 0, 2, 3, 4}))
	assert.Equal(t, 3, MaxChunksToSortedCleverSolution([]int{1, 0, 3, 2, 4}))
	assert.Equal(t, 1, MaxChunksToSortedCleverSolution([]int{4, 3, 2, 1, 0}))
}
