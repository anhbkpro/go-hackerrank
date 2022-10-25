package minimum_window_substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinWindow(t *testing.T) {
	assert.Equal(t, "BANC", MinWindow("ADOBECODEBANC", "ABC"))
	assert.Equal(t, "a", MinWindow("a", "a"))
	assert.Equal(t, "", MinWindow("a", "aa"))
	assert.Equal(t, "aa", MinWindow("aa", "aa"))
}
