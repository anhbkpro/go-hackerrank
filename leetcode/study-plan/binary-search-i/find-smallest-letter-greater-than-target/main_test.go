package find_smallest_letter_greater_than_target

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	assert.Equal(t, byte('j'), NextGreatestLetter([]byte{'c', 'f', 'j'}, 'g'))
	assert.Equal(t, byte('f'), NextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	assert.Equal(t, byte('c'), NextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}
