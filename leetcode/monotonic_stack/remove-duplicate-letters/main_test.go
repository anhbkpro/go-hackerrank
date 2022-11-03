package remove_duplicate_letters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	assert.Equal(t, "abc", RemoveDuplicateLetters("abacb"))
	assert.Equal(t, "eacb", RemoveDuplicateLetters("ecbacba"))
	assert.Equal(t, "adbc", RemoveDuplicateLetters("cdadabcc"))
	assert.Equal(t, "abc", RemoveDuplicateLetters("bcabc"))
	assert.Equal(t, "acdb", RemoveDuplicateLetters("cbacdcbc"))
}
