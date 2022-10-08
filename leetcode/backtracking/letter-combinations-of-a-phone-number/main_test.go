package letter_combinations_of_a_phone_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	assert.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, LetterCombinations("23"))
	assert.Equal(t, []string{}, LetterCombinations(""))
	assert.Equal(t, []string{"a", "b", "c"}, LetterCombinations("2"))
}
