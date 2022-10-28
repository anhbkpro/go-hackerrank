package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	assert.Equal(t, "YES", IsBalanced("{[()]}"))
	assert.Equal(t, "NO", IsBalanced("{[(])}"))
	assert.Equal(t, "YES", IsBalanced("{{[[(())]]}}"))
	assert.Equal(t, "YES", IsBalanced("{{([])}}"))
	assert.Equal(t, "NO", IsBalanced("{{)[](}}"))
	assert.Equal(t, "YES", IsBalanced("{(([])[])[]}"))
	assert.Equal(t, "NO", IsBalanced("{(([])[])[]]}"))
	assert.Equal(t, "YES", IsBalanced("{(([])[])[]}[]"))
}
