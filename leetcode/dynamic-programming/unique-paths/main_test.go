package unique_paths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	assert.Equal(t, 3, UniquePaths(3, 2))
	assert.Equal(t, 28, UniquePaths(3, 7))
}
