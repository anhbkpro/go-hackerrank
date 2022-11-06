package smallest_subsequence_of_distinct_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSmallestSubsequence(t *testing.T) {
	assert.Equal(t, "abc", SmallestSubsequence("bcabc"))
	assert.Equal(t, "acdb", SmallestSubsequence("cbacdcbc"))
}
