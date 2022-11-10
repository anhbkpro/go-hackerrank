package sort_characters_by_frequency

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	// It's ok when this test case can fail, `eert` also acceptable
	assert.Equal(t, "eetr", FrequencySort("tree"))
	assert.Equal(t, "cccaaa", FrequencySort("cccaaa"))
	assert.Equal(t, "bbAa", FrequencySort("Aabb"))
}
