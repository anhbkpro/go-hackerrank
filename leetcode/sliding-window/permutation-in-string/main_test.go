package permutation_in_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	assert.Equal(t, false, CheckInclusion("hello", "ooolleoooleh"))
	assert.Equal(t, true, CheckInclusion("ab", "eidbaooo"))
	assert.Equal(t, false, CheckInclusion("ab", "eidboaoo"))
}
