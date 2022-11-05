package number_of_substrings_with_only_1s

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumSub(t *testing.T) {
	assert.Equal(t, 21, NumSub("111111"))
	assert.Equal(t, 2, NumSub("101"))
	assert.Equal(t, 9, NumSub("0110111"))
}
