package generate_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	got := GenerateParenthesis(3)
	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, got)
}
