package contains_duplicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	got := ContainsDuplicate([]int{1, 2, 3, 1})
	assert.Equal(t, true, got)
}
