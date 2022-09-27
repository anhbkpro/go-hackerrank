package kangaroo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKangaroo(t *testing.T) {
	got := Kangaroo(0, 2, 5, 3)
	assert.Equal(t, "NO", got)

	got2 := Kangaroo(0, 3, 4, 2)
	assert.Equal(t, "YES", got2)
}
