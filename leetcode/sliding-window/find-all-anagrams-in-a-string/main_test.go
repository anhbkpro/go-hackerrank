package find_all_anagrams_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	assert.Equal(t, []int{1}, FindAnagrams("baa", "aa"))
	assert.Equal(t, []int{0, 6}, FindAnagrams("cbaebabacd", "abc"))
	assert.Equal(t, []int{0, 1, 2}, FindAnagrams("abab", "ab"))
}
