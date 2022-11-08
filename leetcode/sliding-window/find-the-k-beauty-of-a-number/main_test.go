package find_the_k_beauty_of_a_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivisorSubstrings(t *testing.T) {
	assert.Equal(t, 2, DivisorSubstrings(240, 2))
	assert.Equal(t, 2, DivisorSubstrings(430043, 2))
}
