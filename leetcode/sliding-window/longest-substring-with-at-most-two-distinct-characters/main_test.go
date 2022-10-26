package longest_substring_with_at_most_two_distinct_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	assert.Equal(t, 3, LengthOfLongestSubstringTwoDistinct("eceba"))
	assert.Equal(t, 5, LengthOfLongestSubstringTwoDistinct("ccaabbb"))
}
