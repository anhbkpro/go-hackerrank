package count_number_of_homogenous_substrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountHomogenous(t *testing.T) {
	assert.Equal(t, 13, CountHomogenous("abbcccaa"))
}
