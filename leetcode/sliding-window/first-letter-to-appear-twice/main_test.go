package first_letter_to_appear_twice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatedCharacter(t *testing.T) {
	assert.Equal(t, []byte("c")[0], RepeatedCharacter("abccbaacz"))
	assert.Equal(t, []byte("d")[0], RepeatedCharacter("abcdd"))
}
