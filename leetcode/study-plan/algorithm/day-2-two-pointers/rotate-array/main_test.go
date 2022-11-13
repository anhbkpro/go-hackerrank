package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(arr, 3)
	assert.Equal(t, []int{5, 6, 7, 1, 2, 3, 4}, arr)
}
