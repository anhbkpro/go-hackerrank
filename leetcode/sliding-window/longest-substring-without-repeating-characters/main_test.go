package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 2, LengthOfLongestSubstringDiscussion("aab"))
	assert.Equal(t, 2, LengthOfLongestSubstringDiscussion("au"))
	assert.Equal(t, 1, LengthOfLongestSubstringDiscussion(" "))
	assert.Equal(t, 3, LengthOfLongestSubstringDiscussion("abcabcbb"))
	assert.Equal(t, 1, LengthOfLongestSubstringDiscussion("bbbbb"))
	assert.Equal(t, 3, LengthOfLongestSubstringDiscussion("pwwkew"))
}
