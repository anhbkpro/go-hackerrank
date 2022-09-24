package count_triplets_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountTriplets(t *testing.T) {
	input := make([]int64, 100)
	for i := range input {
		input[i] = 1
	}
	got := CountTriplets(input, 1)
	assert.Equal(t, int64(161700), got)

	input2 := []int64{1, 3, 9, 9, 27, 81}
	got2 := CountTriplets(input2, 3)
	assert.Equal(t, int64(6), got2)
}
