package chocolate_feast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChocolateFeast(t *testing.T) {
	//got1 := ChocolateFeast(10, 2, 5)
	//got2 := ChocolateFeast(12, 4, 4)
	got3 := ChocolateFeast(6, 2, 2)
	//assert.Equal(t, int32(6), got1)
	//assert.Equal(t, int32(3), got2)
	assert.Equal(t, int32(5), got3)
}
