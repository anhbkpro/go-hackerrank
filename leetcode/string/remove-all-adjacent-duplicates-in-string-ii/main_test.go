package remove_all_adjacent_duplicates_in_string_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, "abcd", RemoveDuplicates("abcd", 2))
	assert.Equal(t, "aa", RemoveDuplicates("deeedbbcccbdaa", 3))
	assert.Equal(t, "ps", RemoveDuplicates("pbbcggttciiippooaais", 2))
}
