package longest_substring_with_at_most_k_distinct_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestKSubstr(t *testing.T) {
	assert.Equal(t, 2, LongestKSubstr("aabbcc", 1))
	assert.Equal(t, 4, LongestKSubstr("aabbcc", 2))
	assert.Equal(t, 6, LongestKSubstr("aabbcc", 3))
	assert.Equal(t, 7, LongestKSubstrV2("cbebebe", 3))
	assert.Equal(t, 5, LongestKSubstrV2("ababcbcca", 2))
	assert.Equal(t, 1, LongestKSubstrV2("abcde", 1))
}
