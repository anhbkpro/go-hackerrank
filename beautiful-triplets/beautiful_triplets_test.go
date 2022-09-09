package beautiful_triplets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBeautifulTriplets(t *testing.T) {
	input := []int32{1, 2, 4, 5, 7, 8, 10}
	got := BeautifulTriplets(3, input)
	assert.Equal(t, int32(3), got)

	input2 := []int32{1, 6, 7, 7, 8, 10, 12, 13, 14, 19}
	got2 := BeautifulTriplets(3, input2)
	assert.Equal(t, int32(2), got2)
}
