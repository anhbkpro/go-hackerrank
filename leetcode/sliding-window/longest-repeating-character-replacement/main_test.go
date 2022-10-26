package longest_repeating_character_replacement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterReplacement(t *testing.T) {
	assert.Equal(t, 4, CharacterReplacement("AAAA", 2))
	assert.Equal(t, 3, CharacterReplacement("AAAB", 0))
	assert.Equal(t, 4, CharacterReplacement("ABAB", 2))
	assert.Equal(t, 4, CharacterReplacement("AABABBA", 1))
}
