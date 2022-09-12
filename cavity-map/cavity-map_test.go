package cavity_map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCavityMap(t *testing.T) {
	input := []string{"1112", "1912", "1892", "1234"}
	got := CavityMap(input)
	assert.Equal(t, []string{"1112", "1X12", "18X2", "1234"}, got)

	input2 := []string{"12", "12"}
	got2 := CavityMap(input2)
	assert.Equal(t, []string{"12", "12"}, got2)
}
