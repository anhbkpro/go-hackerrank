package minimum_deletions_to_make_character_frequencies_unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDeletions(t *testing.T) {
	assert.Equal(t, 2, MinDeletions("aaabbbcc"))
	assert.Equal(t, 2, MinDeletions("ceabaacb"))
	assert.Equal(t, 0, MinDeletions("aab"))
}
