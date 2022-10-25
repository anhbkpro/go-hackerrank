package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 2, LengthOfLongestSubstring("aab"))
	assert.Equal(t, 2, LengthOfLongestSubstring("au"))
	assert.Equal(t, 1, LengthOfLongestSubstring(" "))
	assert.Equal(t, 3, LengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, LengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, LengthOfLongestSubstring("pwwkew"))
}
