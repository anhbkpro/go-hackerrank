package equality_in_a_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualizeArray(t *testing.T) {
	input := []int32{3, 3, 2, 1, 3}
	got := EqualizeArray(input)
	assert.Equal(t, int32(2), got)

	input2 := []int32{1, 2, 3, 1, 2, 3, 3, 3}
	got2 := EqualizeArray(input2)
	assert.Equal(t, int32(4), got2)
}
