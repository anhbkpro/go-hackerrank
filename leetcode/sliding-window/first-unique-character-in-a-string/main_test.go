package first_unique_character_in_a_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	assert.Equal(t, -1, FirstUniqChar("aadadaad"))
	assert.Equal(t, 0, FirstUniqChar("leetcode"))
	assert.Equal(t, 2, FirstUniqChar("loveleetcode"))
	assert.Equal(t, -1, FirstUniqChar("aabb"))
}
