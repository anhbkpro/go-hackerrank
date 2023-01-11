package minimum_deletions_to_make_character_frequencies_unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinDeletions(t *testing.T) {
	assert.Equal(t, 2, MinDeletionsDiscuss("aaabbbcc"))
	assert.Equal(t, 2, MinDeletionsDiscuss("ceabaacb"))
	assert.Equal(t, 0, MinDeletionsDiscuss("aab"))
}
