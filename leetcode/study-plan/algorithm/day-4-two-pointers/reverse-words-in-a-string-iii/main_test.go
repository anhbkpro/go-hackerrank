package reverse_words_in_a_string_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "e", ReverseWords("e"))
	assert.Equal(t, "edoc", ReverseWords("code"))
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", ReverseWords("Let's take LeetCode contest"))
}
