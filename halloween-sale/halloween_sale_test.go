package halloween_sale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHowManyGames(t *testing.T) {
	got := HowManyGames(20, 3, 6, 80)
	assert.Equal(t, int32(6), got)

	got2 := HowManyGames(20, 3, 6, 85)
	assert.Equal(t, int32(7), got2)

	got3 := HowManyGames(20, 3, 6, 70)
	assert.Equal(t, int32(5), got3)
}
