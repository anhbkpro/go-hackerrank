package number_of_substrings_containing_all_three_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfSubstrings(t *testing.T) {
	assert.Equal(t, 10, NumberOfSubstrings("abcabc"))
	assert.Equal(t, 3, NumberOfSubstrings("aaacb"))
}
