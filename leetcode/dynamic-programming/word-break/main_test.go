package word_break

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreak(t *testing.T) {
	assert.Equal(t, true, WordBreak("leetcode", []string{"leet", "code"}))
	assert.Equal(t, true, WordBreak("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, false, WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
