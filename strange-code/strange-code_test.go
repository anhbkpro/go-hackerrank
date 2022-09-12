package strange_code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrangeCounter(t *testing.T) {
	got := StrangeCounter(4)
	assert.Equal(t, int64(6), got)

	got2 := StrangeCounter(17)
	assert.Equal(t, int64(5), got2)
}
