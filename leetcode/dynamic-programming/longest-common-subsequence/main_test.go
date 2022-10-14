package longest_common_subsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert.Equal(t, 3, LongestCommonSubsequence("abcde", "ace"))
	assert.Equal(t, 3, LongestCommonSubsequence("abc", "abc"))
	assert.Equal(t, 0, LongestCommonSubsequence("abc", "def"))
}
