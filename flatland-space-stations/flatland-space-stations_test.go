package flatland_space_stations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatlandSpaceStations(t *testing.T) {
	input := []int32{0, 4}
	got := FlatlandSpaceStations(5, input)
	assert.Equal(t, int32(2), got)

	input2 := []int32{0, 1, 2, 4, 3, 5}
	got2 := FlatlandSpaceStations(6, input2)
	assert.Equal(t, int32(0), got2)
}
