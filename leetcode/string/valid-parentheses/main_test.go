package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, false, IsValid("){"))
	assert.Equal(t, false, IsValid("(("))
	assert.Equal(t, true, IsValid("()"))
	assert.Equal(t, true, IsValid("(){}[]"))
	assert.Equal(t, false, IsValid("(){][}"))
}
