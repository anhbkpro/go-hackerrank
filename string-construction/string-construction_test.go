package string_construction

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringConstruction(t *testing.T) {
	got := StringConstruction("abcd")
	assert.Equal(t, int32(4), got)

	got2 := StringConstruction("abab")
	assert.Equal(t, int32(2), got2)

	got3 := StringConstruction("scfg")
	assert.Equal(t, int32(4), got3)

	got4 := StringConstruction("bccb")
	assert.Equal(t, int32(2), got4)
}
